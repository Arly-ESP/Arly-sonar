package main

import (
	"github.com/arly/arlyApi/config"
	"github.com/arly/arlyApi/database"
	"github.com/arly/arlyApi/index"
	"github.com/arly/arlyApi/utilities"
	"github.com/gofiber/fiber/v2"

	_ "github.com/arly/arlyApi/docs"
)

// @title Arly API
// @version 1.0
// @description This is the Arly API server documentation
// @termsOfService http://swagger.io/terms/

// @contact.name Arly API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5050
// @BasePath /
// @schemes http
func main() {
	utilities.InitializeLogger()

	utilities.LogInfo("Starting Arly API server...")

	database.ConnectDb()
	utilities.LogInfo("Database connection established successfully.")

	app := fiber.New()
	index.PrepareApp(app)

	serverURL := config.ServerUrl()
	if err := app.Listen(serverURL); err != nil {
		utilities.LogFatal("Failed to start the server", err)
	}
}
