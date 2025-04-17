package controllers_test

import (
	"encoding/json"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/arly/arlyApi/controllers"
	"github.com/arly/arlyApi/database"
	"github.com/arly/arlyApi/models"
	tests_db "github.com/arly/arlyApi/tests/db"
	tests_services "github.com/arly/arlyApi/tests/services"
	"github.com/arly/arlyApi/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
)

func setupTestEnv() (*fiber.App, *gorm.DB) {
	utilities.InitializeLogger()
	tests_db.ConnectDb()
	utilities.LogInfo("Database connection established successfully.")

	app := fiber.New()
	mockDB := tests_db.Database.Db

	return app, mockDB
}

func TestRegisterUser_Success(t *testing.T) {
	app, mockDB := setupTestEnv()

	patches := gomonkey.ApplyGlobalVar(&database.Database, database.DbInstance{Db: mockDB})
	defer patches.Reset()

	userData := models.User{
		FirstName: "Test",
		LastName:  "User",
		Email:     "test_arly_5076107@yopmail.com",
		Password:  "password123",
	}

	jsonData, _ := json.Marshal(userData)

	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	ctx.Request().SetBody(jsonData)
	ctx.Request().Header.SetContentType("application/json")

	emailService := tests_services.NewEmailService()
	ctx.Locals("emailService", emailService)

	err := controllers.RegisterUser(ctx)
	assert.NoError(t, err)

	assert.Equal(t, fiber.StatusCreated, ctx.Response().StatusCode())

	var respBody map[string]interface{}
	json.Unmarshal(ctx.Response().Body(), &respBody)

	// TODO:
	// assert that a new user has been successfully created by making a transaction to the database
}
