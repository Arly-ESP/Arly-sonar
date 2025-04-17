package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/arly/arlyApi/utilities"
	"fmt"
	"github.com/fatih/color"
)

func TimeSpentMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		startTime := time.Now()

		err := c.Next()
		timeSpent := time.Since(startTime)

		c.Set("X-Time-Spent", timeSpent.String())
		c.Set("X-Date", time.Now().Format(time.RFC3339))

		return err
	}
}

func RequestLogger(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()

	duration := time.Since(start)
	statusCode := c.Response().StatusCode()

	var logColor func(a ...interface{}) string
	switch {
	case statusCode >= 500:
		logColor = color.New(color.FgRed).SprintFunc() 
	case statusCode >= 400:
		logColor = color.New(color.FgYellow).SprintFunc() 
	case statusCode >= 300:
		logColor = color.New(color.FgCyan).SprintFunc()
	case statusCode == 200:
		logColor = color.New(color.FgGreen).SprintFunc() 
	default:
		logColor = color.New(color.FgWhite).SprintFunc() 
	}

	logMessage := fmt.Sprintf(
		"Path: %s, Method: %s, Status Code: %d, Duration: %s",
		c.Path(), c.Method(), statusCode, duration.String(),
	)

	if err != nil {
		utilities.LogError(logColor(logMessage), err)
	} else {
		utilities.LogInfo(logColor(logMessage))
	}

	return err
}
