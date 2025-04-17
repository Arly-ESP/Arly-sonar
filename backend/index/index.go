package index

import (
	"time"

	"github.com/arly/arlyApi/routes"
	"github.com/arly/arlyApi/services"
	"github.com/arly/arlyApi/utilities"
	"github.com/arly/arlyApi/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// healthCheck godoc
// @Summary Health check endpoint
// @Description Returns a JSON response with a welcome message, the current time, and the status of the API.
// @Tags Health
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{} "Success"
// @Router /api/health [get]
func healthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome to the API",
		"time":    time.Now().Format(time.RFC3339),
		"status":  "up",
	})
}

func PrepareApp(app *fiber.App) {

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", 
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(middleware.TimeSpentMiddleware())
	app.Use(middleware.RequestLogger)


	emailService := services.NewEmailService()

	if err := emailService.PingSMTP(); err != nil {
		utilities.LogFatal("Failed to connect to SMTP server", err)
	}

	// send test email
	// subject := "Test email from Arly API"
	// body := "This is a test email from the Arly API server. If you received this email, the SMTP server is working correctly."
	// to := "contact@sanlamamba.com"
	// if err := emailService.SendEmail([]string{to}, subject, body); err != nil {
	// 	utilities.LogFatal("Failed to send test email", err)
	// }


	utilities.LogInfo("SMTP server connection established successfully.")

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("emailService", emailService)
		return c.Next()
	})

	SetupRoutes(app)
}

func SetupRoutes(app *fiber.App) {

	app.Get("/api/health", healthCheck)
	app.Get("/swagger/*", swagger.HandlerDefault)

	routes.UserRoutes(app)
	routes.UserAuth(app)
	routes.ChatRoutes(app)
	routes.SurveyRoutes(app)
}
