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

func TestSurveyRoutes(t *testing.T) {
	app := fiber.New()

	// Patch JWTMiddleware and AdminMiddleware to simply call the next handler.
	jwtPatch := gomonkey.ApplyFunc(middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		return c.Next()
	})
	defer jwtPatch.Reset()

	adminPatch := gomonkey.ApplyFunc(middleware.AdminMiddleware, func(c *fiber.Ctx) error {
		return c.Next()
	})
	defer adminPatch.Reset()

	// Patch controller functions to return dummy JSON responses.
	createSurveyPatch := gomonkey.ApplyFunc(controllers.CreateSurvey, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "CreateSurvey called"})
	})
	defer createSurveyPatch.Reset()

	updateSurveyPatch := gomonkey.ApplyFunc(controllers.UpdateSurvey, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "UpdateSurvey called"})
	})
	defer updateSurveyPatch.Reset()

	deleteSurveyPatch := gomonkey.ApplyFunc(controllers.DeleteSurvey, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "DeleteSurvey called"})
	})
	defer deleteSurveyPatch.Reset()

	getAllSurveysPatch := gomonkey.ApplyFunc(controllers.GetAllSurveys, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "GetAllSurveys called"})
	})
	defer getAllSurveysPatch.Reset()

	getSurveyDetailsPatch := gomonkey.ApplyFunc(controllers.GetSurveyDetails, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "GetSurveyDetails called"})
	})
	defer getSurveyDetailsPatch.Reset()

	getSurveyBySlugPatch := gomonkey.ApplyFunc(controllers.GetSurveyBySlug, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "GetSurveyBySlug called"})
	})
	defer getSurveyBySlugPatch.Reset()

	submitSurveyResponsePatch := gomonkey.ApplyFunc(controllers.SubmitSurveyResponse, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "SubmitSurveyResponse called"})
	})
	defer submitSurveyResponsePatch.Reset()

	getAllSurveyResponsesPatch := gomonkey.ApplyFunc(controllers.GetAllSurveyResponses, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "GetAllSurveyResponses called"})
	})
	defer getAllSurveyResponsesPatch.Reset()

	getUserSurveyResponsePatch := gomonkey.ApplyFunc(controllers.GetUserSurveyResponse, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "GetUserSurveyResponse called"})
	})
	defer getUserSurveyResponsePatch.Reset()

	// Register the survey routes.
	SurveyRoutes(app)

	// Helper function to send a request and verify the JSON response message.
	testEndpoint := func(method, path, expectedMsg string) {
		req := httptest.NewRequest(method, path, nil)
		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		var body map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&body)
		assert.NoError(t, err)
		assert.Equal(t, expectedMsg, body["message"])
	}

	// Survey Management (JWT + Admin protected)
	testEndpoint("POST", "/api/surveys", "CreateSurvey called")
	testEndpoint("PUT", "/api/surveys/123", "UpdateSurvey called")
	testEndpoint("DELETE", "/api/surveys/123", "DeleteSurvey called")

	// Public Survey Routes
	testEndpoint("GET", "/api/surveys", "GetAllSurveys called")
	testEndpoint("GET", "/api/surveys/123", "GetSurveyDetails called")

	// Get survey by slug (JWT protected)
	testEndpoint("GET", "/api/surveys/slug/test-slug", "GetSurveyBySlug called")

	// User Survey Responses
	testEndpoint("POST", "/api/surveys/123/responses", "SubmitSurveyResponse called")
	testEndpoint("GET", "/api/surveys/123/responses", "GetAllSurveyResponses called")
	testEndpoint("GET", "/api/surveys/123/responses/456", "GetUserSurveyResponse called")
}
