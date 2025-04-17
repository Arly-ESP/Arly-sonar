package controllers

import (
	"errors"
	"strconv"
	"time"

	"github.com/arly/arlyApi/database"
	"github.com/arly/arlyApi/models"
	"github.com/arly/arlyApi/services"
	"github.com/arly/arlyApi/utilities"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func getUserIDFromContext(c *fiber.Ctx) (uint, error) {
	userID := c.Locals("userID")
	if userID == nil {
		utilities.LogError("Unauthorized access: userID not found in context", nil)
		return 0, errors.New("unauthorized access: userID not found in context")
	}
	utilities.LogInfo("Retrieved userID from context")
	return userID.(uint), nil
}

func UpdateUserActivity(userID uint) error {
	currentDate := time.Now().Truncate(24 * time.Hour)

	var activity models.UserActivity
	err := database.Database.Db.Where("user_id = ? AND date = ?", userID, currentDate).First(&activity).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		activity = models.UserActivity{
			UserID:      userID,
			Date:        currentDate,
			Mood:        "neutral",
			MessageCount: 1,
		}
		if createErr := database.Database.Db.Create(&activity).Error; createErr != nil {
			utilities.LogError("Failed to create user activity", createErr)
			return errors.New("failed to create user activity")
		}
		utilities.LogInfo("User activity created successfully")
	} else if err != nil {
		utilities.LogError("Failed to fetch user activity", err)
		return errors.New("failed to fetch user activity")
	} else {
		activity.MessageCount++
		if saveErr := database.Database.Db.Save(&activity).Error; saveErr != nil {
			utilities.LogError("Failed to update user activity", saveErr)
			return errors.New("failed to update user activity")
		}
		utilities.LogInfo("User activity updated successfully")
	}

	return nil
}

// ChatWithAI godoc
// @Summary Interact with the AI assistant
// @Description Allows an authenticated user to send a message to the AI assistant. The user's activity is tracked and their message is stored. Returns the AI's response.
// @Tags Chat
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer #)
// @Param chatReq body ChatRequest true "Message to send to the AI assistant"
// @Success 200 {object} ChatResponse "AI response with message and chat details"
// @Failure 400 {object} ErrorResponse "Invalid request payload"
// @Failure 401 {object} ErrorResponse "Unauthorized access, token missing or invalid"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/chat [post]
func ChatWithAI(c *fiber.Ctx) error {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	var chatReq ChatRequest
	if err := c.BodyParser(&chatReq); err != nil {
		utilities.LogError("Invalid request payload for ChatWithAI", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	aiResponse, chatID, messageID, err := services.HandleChat(userID, chatReq.Message)
	if err != nil {
		utilities.LogError("Failed to handle chat", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := UpdateUserActivity(userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user activity", "details": err.Error()})
	}

	utilities.LogInfo("ChatWithAI executed successfully")
	return c.Status(fiber.StatusOK).JSON(ChatResponse{
		Response:   aiResponse,
		MessageID:  messageID,
		ChatID:     chatID,
		Timestamp:  time.Now(),
	})
}

// GetChatsByUserID godoc
// @Summary Retrieve chats and messages for the authenticated user
// @Description Fetches all chats and their messages for the authenticated user. Returns a list of chat IDs with their associated messages.
// @Tags Chat
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer #)
// @Success 200 {array} ChatWithMessages "List of chats with their messages"
// @Failure 401 {object} ErrorResponse "Unauthorized access, token missing or invalid"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/user/chats [get]
func GetChatsByUserID(c *fiber.Ctx) error {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	var chats []models.Chat
	if err := database.Database.Db.Where("user_id = ?", userID).Find(&chats).Error; err != nil {
		utilities.LogError("Failed to fetch chats for user", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch chats"})
	}

	if len(chats) == 0 {
		utilities.LogInfo("No chats found for user")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No chats found for this user"})
	}

	var chatsWithMessages []ChatWithMessages
	for _, chat := range chats {
		var messages []models.Message
		if err := database.Database.Db.Where("chat_id = ?", chat.ID).Find(&messages).Error; err != nil {
			utilities.LogError("Failed to fetch messages for chat", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch messages for a chat"})
		}

		chatsWithMessages = append(chatsWithMessages, ChatWithMessages{
			ChatID:   chat.ID,
			Messages: messages,
		})
	}

	utilities.LogInfo("GetChatsByUserID executed successfully")
	return c.Status(fiber.StatusOK).JSON(chatsWithMessages)
}

// GetChatMessages godoc
// @Summary Fetch messages for a specific chat
// @Description Allows an authenticated user to retrieve all messages for a specific chat by providing the chat ID.
// @Tags Chat
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer #)
// @Param chat_id path uint true "Chat ID to retrieve messages for"
// @Success 200 {array} models.Message "List of messages for the specified chat"
// @Failure 400 {object} ErrorResponse "Invalid or missing Chat ID"
// @Failure 401 {object} ErrorResponse "Unauthorized access, token missing or invalid"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/chat/{chat_id}/messages [get]
func GetChatMessages(c *fiber.Ctx) error {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	chatIDParam := c.Params("chat_id")
	chatID, err := strconv.ParseUint(chatIDParam, 10, 64)
	if err != nil || chatID == 0 {
		utilities.LogError("Invalid chat ID provided", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid chat ID"})
	}

	var chat models.Chat
	if err := database.Database.Db.Where("id = ? AND user_id = ?", chatID, userID).First(&chat).Error; err != nil {
		utilities.LogError("Unauthorized access to chat", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized access to this chat"})
	}

	var messages []models.Message
	if err := database.Database.Db.Where("chat_id = ?", chatID).Find(&messages).Error; err != nil {
		utilities.LogError("Failed to fetch messages for chat", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch messages"})
	}

	if len(messages) == 0 {
		utilities.LogInfo("No messages found for chat")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No messages found for this chat"})
	}

	utilities.LogInfo("GetChatMessages executed successfully")
	return c.Status(fiber.StatusOK).JSON(messages)
}
