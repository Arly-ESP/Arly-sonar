package middleware_test

import (
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/arly/arlyApi/middleware"
	"github.com/arly/arlyApi/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestTimeSpentMiddleware(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	ctx.Request().Header.SetContentType("application/json")

	patchNext := gomonkey.ApplyMethod(ctx, "Next", func(c *fiber.Ctx) error {
		return nil
	})
	defer patchNext.Reset()

	handler := middleware.TimeSpentMiddleware()
	err := handler(ctx)
	assert.NoError(t, err)

	xTimeSpent := ctx.Response().Header.Peek("X-Time-Spent")
	assert.NotEmpty(t, string(xTimeSpent), "X-Time-Spent header should be set")

	xDate := ctx.Response().Header.Peek("X-Date")
	assert.NotEmpty(t, string(xDate), "X-Date header should be set")
	_, parseErr := time.Parse(time.RFC3339, string(xDate))
	assert.NoError(t, parseErr, "X-Date header should be a valid RFC3339 timestamp")
}

func TestRequestLogger_Success(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	ctx.Request().SetRequestURI("/test")
	ctx.Method("GET")
	ctx.Request().Header.SetContentType("application/json")

	patchNext := gomonkey.ApplyMethod(ctx, "Next", func(c *fiber.Ctx) error {
		c.Status(fiber.StatusOK)
		return nil
	})
	defer patchNext.Reset()

	patchLogInfo := gomonkey.ApplyFunc(utilities.LogInfo, func(msg string) {})
	defer patchLogInfo.Reset()
	patchLogError := gomonkey.ApplyFunc(utilities.LogError, func(msg string, err error) {})
	defer patchLogError.Reset()

	err := middleware.RequestLogger(ctx)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, ctx.Response().StatusCode())
}

func TestRequestLogger_Error(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	ctx.Request().SetRequestURI("/error")
	ctx.Method("GET")
	ctx.Request().Header.SetContentType("application/json")

	simulatedError := fiber.NewError(fiber.StatusInternalServerError, "simulated error")
	patchNext := gomonkey.ApplyMethod(ctx, "Next", func(c *fiber.Ctx) error {
		c.Status(fiber.StatusInternalServerError)
		return simulatedError
	})
	defer patchNext.Reset()

	patchLogInfo := gomonkey.ApplyFunc(utilities.LogInfo, func(msg string) {})
	defer patchLogInfo.Reset()
	patchLogError := gomonkey.ApplyFunc(utilities.LogError, func(msg string, err error) {})
	defer patchLogError.Reset()

	err := middleware.RequestLogger(ctx)
	assert.Equal(t, simulatedError, err)
	assert.Equal(t, fiber.StatusInternalServerError, ctx.Response().StatusCode())
}
