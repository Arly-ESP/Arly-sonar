package routes

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/arly/arlyApi/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// TestUserAuthRoutes verifies that the UserAuth function properly registers
// the authentication endpoints and that the patched controller functions are called.
func TestUserAuthRoutes(t *testing.T) {
	app := fiber.New()

	// Patch controller functions to return known JSON responses.
	patchRegister := gomonkey.ApplyFunc(controllers.RegisterUser, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "register called"})
	})
	defer patchRegister.Reset()

	patchLogin := gomonkey.ApplyFunc(controllers.LoginUser, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "login called"})
	})
	defer patchLogin.Reset()

	patchVerify := gomonkey.ApplyFunc(controllers.VerifyUser, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "verify called"})
	})
	defer patchVerify.Reset()

	patchRequestReset := gomonkey.ApplyFunc(controllers.RequestPasswordReset, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "password reset request called"})
	})
	defer patchRequestReset.Reset()

	patchResetPassword := gomonkey.ApplyFunc(controllers.ResetPassword, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "password reset called"})
	})
	defer patchResetPassword.Reset()

	// Register the authentication routes.
	UserAuth(app)

	// Test POST /api/register
	reqRegister := httptest.NewRequest("POST", "/api/register", nil)
	respRegister, err := app.Test(reqRegister, -1)
	assert.NoError(t, err)
	var bodyRegister map[string]interface{}
	json.NewDecoder(respRegister.Body).Decode(&bodyRegister)
	assert.Equal(t, "register called", bodyRegister["message"])

	// Test POST /api/login
	reqLogin := httptest.NewRequest("POST", "/api/login", nil)
	respLogin, err := app.Test(reqLogin, -1)
	assert.NoError(t, err)
	var bodyLogin map[string]interface{}
	json.NewDecoder(respLogin.Body).Decode(&bodyLogin)
	assert.Equal(t, "login called", bodyLogin["message"])

	// Test POST /api/verify
	reqVerify := httptest.NewRequest("POST", "/api/verify", nil)
	respVerify, err := app.Test(reqVerify, -1)
	assert.NoError(t, err)
	var bodyVerify map[string]interface{}
	json.NewDecoder(respVerify.Body).Decode(&bodyVerify)
	assert.Equal(t, "verify called", bodyVerify["message"])

	// Test GET /api/password-reset
	reqPasswordResetGet := httptest.NewRequest("GET", "/api/password-reset", nil)
	respPasswordResetGet, err := app.Test(reqPasswordResetGet, -1)
	assert.NoError(t, err)
	var bodyPasswordResetGet map[string]interface{}
	json.NewDecoder(respPasswordResetGet.Body).Decode(&bodyPasswordResetGet)
	assert.Equal(t, "password reset request called", bodyPasswordResetGet["message"])

	// Test POST /api/password-reset
	reqPasswordResetPost := httptest.NewRequest("POST", "/api/password-reset", nil)
	respPasswordResetPost, err := app.Test(reqPasswordResetPost, -1)
	assert.NoError(t, err)
	var bodyPasswordResetPost map[string]interface{}
	json.NewDecoder(respPasswordResetPost.Body).Decode(&bodyPasswordResetPost)
	assert.Equal(t, "password reset called", bodyPasswordResetPost["message"])
}
