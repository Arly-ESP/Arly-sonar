package routes

import (
	"github.com/arly/arlyApi/controllers"
	"github.com/arly/arlyApi/middleware"
	"github.com/gofiber/fiber/v2"
)

func ChatRoutes(app *fiber.App) {
	app.Post("/api/chat", middleware.JWTMiddleware, controllers.ChatWithAI)
	app.Get("/api/chat/:chat_id/messages", middleware.JWTMiddleware, controllers.GetChatMessages)
	app.Get("/api/user/chats", middleware.JWTMiddleware, controllers.GetChatsByUserID)
}
