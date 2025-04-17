package routes

import (
	"github.com/arly/arlyApi/controllers"
	"github.com/arly/arlyApi/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	// Protected routes for authenticated users
	app.Get("/api/user", middleware.JWTMiddleware, controllers.GetUser)
	app.Put("/app/user", middleware.JWTMiddleware, controllers.UpdateUser)
	app.Delete("/api/user", middleware.JWTMiddleware, controllers.DeleteCurrentUser)

	// Activity routes (protected for authenticated users)
	app.Post("/api/mood", middleware.JWTMiddleware, controllers.LogUserMood)
	app.Get("/api/user/activity", middleware.JWTMiddleware, controllers.GetUserActivity)
	app.Get("/api/user/activities", middleware.JWTMiddleware, controllers.GetUserActivities)

	// Admin-only routes 
	app.Get("/api/admin/users", controllers.GetUsers)  
	app.Get("/api/admin/users/:id", controllers.GetUserDetails) 	
	app.Delete("/api/admin/users/:id", controllers.DeleteUser)
 }
