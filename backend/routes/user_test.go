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

func TestUserRoutes(t *testing.T) {
	app := fiber.New()

	// Patch the JWT middleware to simply call the next handler.
	patchJWT := gomonkey.ApplyFunc(middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		return c.Next()
	})
	defer patchJWT.Reset()

	// Patch the controllers to return dummy responses.
	patchGetUser := gomonkey.ApplyFunc(controllers.GetUser, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "GetUser called"})
	})
	defer patchGetUser.Reset()

	patchUpdateUser := gomonkey.ApplyFunc(controllers.UpdateUser, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "UpdateUser called"})
	})
	defer patchUpdateUser.Reset()

	patchDeleteCurrentUser := gomonkey.ApplyFunc(controllers.DeleteCurrentUser, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "DeleteCurrentUser called"})
	})
	defer patchDeleteCurrentUser.Reset()

	patchLogUserMood := gomonkey.ApplyFunc(controllers.LogUserMood, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "LogUserMood called"})
	})
	defer patchLogUserMood.Reset()

	patchGetUserActivity := gomonkey.ApplyFunc(controllers.GetUserActivity, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "GetUserActivity called"})
	})
	defer patchGetUserActivity.Reset()

	patchGetUserActivities := gomonkey.ApplyFunc(controllers.GetUserActivities, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "GetUserActivities called"})
	})
	defer patchGetUserActivities.Reset()

	patchGetUsers := gomonkey.ApplyFunc(controllers.GetUsers, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "GetUsers called"})
	})
	defer patchGetUsers.Reset()

	patchGetUserDetails := gomonkey.ApplyFunc(controllers.GetUserDetails, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "GetUserDetails called"})
	})
	defer patchGetUserDetails.Reset()

	patchDeleteUser := gomonkey.ApplyFunc(controllers.DeleteUser, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "DeleteUser called"})
	})
	defer patchDeleteUser.Reset()

	// Register the user routes.
	UserRoutes(app)

	// Create a helper function to send a request and verify response.
	testEndpoint := func(method, path, expectedMsg string) {
		req := httptest.NewRequest(method, path, nil)
		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		var body map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&body)
		assert.NoError(t, err)
		assert.Equal(t, expectedMsg, body["message"])
	}

	// Protected routes for authenticated users:
	testEndpoint("GET", "/api/user", "GetUser called")
	testEndpoint("PUT", "/app/user", "UpdateUser called")
	testEndpoint("DELETE", "/api/user", "DeleteCurrentUser called")

	// Activity routes (protected for authenticated users):
	testEndpoint("POST", "/api/mood", "LogUserMood called")
	testEndpoint("GET", "/api/user/activity", "GetUserActivity called")
	testEndpoint("GET", "/api/user/activities", "GetUserActivities called")

	// Admin-only routes:
	testEndpoint("GET", "/api/admin/users", "GetUsers called")
	testEndpoint("GET", "/api/admin/users/1", "GetUserDetails called")
	testEndpoint("DELETE", "/api/admin/users/1", "DeleteUser called")
}
