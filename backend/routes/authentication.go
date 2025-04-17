package routes

import (
	"github.com/arly/arlyApi/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserAuth(app *fiber.App) {
	app.Post("/api/register", controllers.RegisterUser)   
	app.Post("/api/login", controllers.LoginUser)        
	app.Post("/api/verify", controllers.VerifyUser)       

	app.Get("/api/password-reset", controllers.RequestPasswordReset)
	app.Post("/api/password-reset", controllers.ResetPassword)      
}
