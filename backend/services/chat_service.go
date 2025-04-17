package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"gorm.io/gorm" 
	"errors"

	"github.com/arly/arlyApi/database"
	"github.com/arly/arlyApi/models"
	"github.com/arly/arlyApi/config"
)


type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIChatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

type ConversationManager struct {
	ChatID   uint
	UserID   uint
	Messages []Message
}

func NewConversationManager(chatID, userID uint) *ConversationManager {
	return &ConversationManager{
		ChatID: chatID,
		UserID: userID,
		Messages: []Message{
			{
				Role: "system",
				Content: "Vous êtes un assistant empathique et bienveillant, formé pour aider les utilisateurs à gérer et à comprendre leurs émotions. " +
					"Votre rôle est de fournir un soutien émotionnel, d'encourager l'utilisateur à exprimer ses sentiments, et de l'aider à trouver des perspectives apaisantes. " +
					"Vous n'êtes ni thérapeute ni professionnel de la santé, et vous ne pouvez pas fournir de conseils médicaux ou diagnostiques. " +
					"Veuillez toujours rester dans votre rôle d'accompagnateur bienveillant, poser des questions pour comprendre et encourager l'utilisateur, " +
					"et éviter tout jugement. En toute circonstance, soyez compréhensif et apportez une perspective positive et encourageante. " +
					"En cas de situation de crise, encouragez discrètement l'utilisateur à consulter un professionnel de santé. " +
					"Répondez toujours en français, soyez empathique et bienveillant, et ne donnez pas de conseils médicaux ou diagnostiques.",
			},
		},
	}
}

func (cm *ConversationManager) AddUserMessage(content string) error {
	message := models.Message{
		Date:         time.Now(),
		MessageType:  "user",
		IsBotMessage: false,
		Content:      content,
		ChatID:       cm.ChatID,
		UserID:       cm.UserID,
	}
	if err := database.Database.Db.Create(&message).Error; err != nil {
		return fmt.Errorf("failed to save user message: %w", err)
	}
	cm.Messages = append(cm.Messages, Message{Role: "user", Content: content})
	return nil
}

func (cm *ConversationManager) AddAssistantMessage(content string, responseTime int) error {
	message := models.Message{
		Date:         time.Now(),
		MessageType:  "assistant",
		IsBotMessage: true,
		Content:      content,
		ResponseTime: responseTime,
		ChatID:       cm.ChatID,
		UserID:       cm.UserID,
	}
	if err := database.Database.Db.Create(&message).Error; err != nil {
		return fmt.Errorf("failed to save assistant message: %w", err)
	}
	cm.Messages = append(cm.Messages, Message{Role: "assistant", Content: content})
	return nil
}

func (cm *ConversationManager) QueryOpenAI() (string, error) {
	requestData := map[string]interface{}{
		"model":      config.GetOpenAiConfig().Model,
		"messages":    cm.Messages,
		"temperature": config.GetOpenAiConfig().Temperature,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request data: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create new request: %w", err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.GetOpenAiConfig().APIKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request to OpenAI failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("OpenAI API error (status %d): %s", resp.StatusCode, string(body))
	}

	var openAIResponse OpenAIChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&openAIResponse); err != nil {
		return "", fmt.Errorf("failed to decode OpenAI response: %w", err)
	}

	if len(openAIResponse.Choices) > 0 {
		return openAIResponse.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no choices returned in OpenAI response")
}

func newChat(userID uint) (models.Chat, error) {
	var context models.Context
	if err := database.Database.Db.Where("user_id = ?", userID).First(&context).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create a new context if none exists
			context = models.Context{UserID: userID}
			if err := database.Database.Db.Create(&context).Error; err != nil {
				return models.Chat{}, fmt.Errorf("failed to create context: %w", err)
			}
		} else {
			return models.Chat{}, fmt.Errorf("failed to fetch context: %w", err)
		}
	}

	var chat models.Chat
	if err := database.Database.Db.Where("user_id = ? AND context_id = ?", userID, context.ID).First(&chat).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create a new chat if none exists
			chat = models.Chat{UserID: userID, ContextID: context.ID}
			if err := database.Database.Db.Create(&chat).Error; err != nil {
				return models.Chat{}, fmt.Errorf("failed to create chat: %w", err)
			}
		} else {
			return models.Chat{}, fmt.Errorf("failed to fetch chat: %w", err)
		}
	}

	return chat, nil
}

func HandleChat(userID uint, userMessage string) (string, uint, uint, error) {
	// Step 1: Get or create a chat
	chat, err := newChat(userID)
	if err != nil {
		return "", 0, 0, fmt.Errorf("failed to create or fetch chat: %w", err)
	}

	// Initialize the conversation manager
	manager := NewConversationManager(chat.ID, userID)

	// Step 2: Add the user's message
	if err := manager.AddUserMessage(userMessage); err != nil {
		return "", 0, 0, fmt.Errorf("failed to add user message: %w", err)
	}

	// Step 3: Query OpenAI for a response
	startTime := time.Now()
	aiResponse, err := manager.QueryOpenAI()
	if err != nil {
		return "", 0, 0, fmt.Errorf("failed to query OpenAI: %w", err)
	}
	responseTime := int(time.Since(startTime).Milliseconds())

	// Step 4: Add the assistant's message
	if err := manager.AddAssistantMessage(aiResponse, responseTime); err != nil {
		return "", 0, 0, fmt.Errorf("failed to add assistant message: %w", err)
	}

	return aiResponse, chat.ID, manager.ChatID, nil
}

func GetChatMessages(chatID string) ([]models.Message, error) {
	var messages []models.Message
	if err := database.Database.Db.Where("chat_id = ?", chatID).Find(&messages).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch messages: %w", err)
	}
	return messages, nil
}