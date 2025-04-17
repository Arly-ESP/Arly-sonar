package routes

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/arly/arlyApi/controllers"
	"github.com/arly/arlyApi/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestChatRoutes(t *testing.T) {
	// Patch JWTMiddleware so that it just calls c.Next()
	patchJWT := gomonkey.ApplyFunc(middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		return c.Next()
	})
	defer patchJWT.Reset()

	// Patch controllers.ChatWithAI to return a dummy JSON response.
	patchChatWithAI := gomonkey.ApplyFunc(controllers.ChatWithAI, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "ChatWithAI called"})
	})
	defer patchChatWithAI.Reset()

	// Patch controllers.GetChatMessages to return a dummy JSON response.
	patchGetChatMessages := gomonkey.ApplyFunc(controllers.GetChatMessages, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "GetChatMessages called"})
	})
	defer patchGetChatMessages.Reset()

	// Patch controllers.GetChatsByUserID to return a dummy JSON response.
	patchGetChatsByUserID := gomonkey.ApplyFunc(controllers.GetChatsByUserID, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "GetChatsByUserID called"})
	})
	defer patchGetChatsByUserID.Reset()

	// Create a new Fiber app and register the ChatRoutes.
	app := fiber.New()
	ChatRoutes(app)

	// Test POST /api/chat
	reqChat := httptest.NewRequest("POST", "/api/chat", nil)
	respChat, err := app.Test(reqChat, -1)
	assert.NoError(t, err)
	assert.Equal(t, 200, respChat.StatusCode)

	var bodyChat map[string]interface{}
	err = json.NewDecoder(respChat.Body).Decode(&bodyChat)
	assert.NoError(t, err)
	assert.Equal(t, "ChatWithAI called", bodyChat["message"])

	// Test GET /api/chat/:chat_id/messages (e.g., chat_id=123)
	reqChatMessages := httptest.NewRequest("GET", "/api/chat/123/messages", nil)
	respChatMessages, err := app.Test(reqChatMessages, -1)
	assert.NoError(t, err)
	assert.Equal(t, 200, respChatMessages.StatusCode)

	var bodyChatMessages map[string]interface{}
	err = json.NewDecoder(respChatMessages.Body).Decode(&bodyChatMessages)
	assert.NoError(t, err)
	assert.Equal(t, "GetChatMessages called", bodyChatMessages["message"])

	// Test GET /api/user/chats
	reqUserChats := httptest.NewRequest("GET", "/api/user/chats", nil)
	respUserChats, err := app.Test(reqUserChats, -1)
	assert.NoError(t, err)
	assert.Equal(t, 200, respUserChats.StatusCode)

	var bodyUserChats map[string]interface{}
	err = json.NewDecoder(respUserChats.Body).Decode(&bodyUserChats)
	assert.NoError(t, err)
	assert.Equal(t, "GetChatsByUserID called", bodyUserChats["message"])
}
