package middleware_test

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/arly/arlyApi/config"
	"github.com/arly/arlyApi/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

// TestJWTMiddleware_NoAuthHeader verifies that if no Authorization header is provided, JWTMiddleware returns 401.
func TestJWTMiddleware_NoAuthHeader(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	ctx.Request().Header.SetContentType("application/json")

	err := middleware.JWTMiddleware(ctx)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, ctx.Response().StatusCode())

	var resp map[string]interface{}
	_ = json.Unmarshal(ctx.Response().Body(), &resp)
	assert.Equal(t, "Your session is invalid. Please log in again.", resp["error"])
}

// TestJWTMiddleware_InvalidFormat verifies that an incorrectly formatted Authorization header returns 401.
func TestJWTMiddleware_InvalidFormat(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	// Set an invalid header (not "Bearer <token>")
	ctx.Request().Header.Set("Authorization", "InvalidTokenFormat")
	ctx.Request().Header.SetContentType("application/json")

	err := middleware.JWTMiddleware(ctx)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, ctx.Response().StatusCode())

	var resp map[string]interface{}
	_ = json.Unmarshal(ctx.Response().Body(), &resp)
	assert.Equal(t, "Your session is invalid. Please log in and try again.", resp["error"])
}

// TestJWTMiddleware_InvalidToken verifies that an invalid token results in a 401.
func TestJWTMiddleware_InvalidToken(t *testing.T) {
	patchToken := gomonkey.ApplyFunc(config.ValidateToken, func(token string) (uint, error) {
		return 0, errors.New("token is invalid")
	})
	defer patchToken.Reset()

	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	ctx.Request().Header.Set("Authorization", "Bearer invalidtoken")
	ctx.Request().Header.SetContentType("application/json")

	err := middleware.JWTMiddleware(ctx)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, ctx.Response().StatusCode())

	var resp map[string]interface{}
	_ = json.Unmarshal(ctx.Response().Body(), &resp)
	assert.Equal(t, "Invalid or expired token", resp["error"])
	assert.True(t, strings.Contains(resp["details"].(string), "token is invalid"))
}

// TestJWTMiddleware_ValidToken verifies that a valid token sets the userID in locals.
func TestJWTMiddleware_ValidToken(t *testing.T) {
	patchToken := gomonkey.ApplyFunc(config.ValidateToken, func(token string) (uint, error) {
		return 2, nil
	})
	defer patchToken.Reset()

	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	ctx.Request().Header.Set("Authorization", "Bearer validtoken")
	ctx.Request().Header.SetContentType("application/json")

	// Patch the Next() method to simulate a next handler.
	patchNext := gomonkey.ApplyMethod(ctx, "Next", func(c *fiber.Ctx) error {
		return nil
	})
	defer patchNext.Reset()

	err := middleware.JWTMiddleware(ctx)
	assert.NoError(t, err)
	userID := ctx.Locals("userID")
	assert.Equal(t, uint(2), userID)
}

// TestAdminMiddleware_NoUserID verifies that if no userID is present, it returns 401.
func TestAdminMiddleware_NoUserID(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	ctx.Request().Header.SetContentType("application/json")
	err := middleware.AdminMiddleware(ctx)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, ctx.Response().StatusCode())

	var resp map[string]interface{}
	_ = json.Unmarshal(ctx.Response().Body(), &resp)
	assert.Equal(t, "Forbidden: Admin access was not granted", resp["error"])
}

// TestAdminMiddleware_NonAdmin verifies that if userID is not 1, it returns 403.
func TestAdminMiddleware_NonAdmin(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	ctx.Request().Header.SetContentType("application/json")
	ctx.Locals("userID", uint(2))
	err := middleware.AdminMiddleware(ctx)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusForbidden, ctx.Response().StatusCode())

	var resp map[string]interface{}
	_ = json.Unmarshal(ctx.Response().Body(), &resp)
	assert.Equal(t, "Forbidden: Admin access required", resp["error"])
}

// TestAdminMiddleware_Admin verifies that if userID is 1, the middleware calls Next().
// We patch the Next() method to avoid nil pointer dereference.
func TestAdminMiddleware_Admin(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	ctx.Request().Header.SetContentType("application/json")
	ctx.Locals("userID", uint(1))

	patchNext := gomonkey.ApplyMethod(ctx, "Next", func(c *fiber.Ctx) error {
		return nil
	})
	defer patchNext.Reset()

	err := middleware.AdminMiddleware(ctx)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, ctx.Response().StatusCode())
}
