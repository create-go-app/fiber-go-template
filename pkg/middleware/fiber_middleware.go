package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// FiberMiddleware provide Fiber's built-in middlewares.
func FiberMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(),
		// Add simple logger.
		logger.New(
			logger.Config{
				Format:     "${time} [${status}] ${method} ${path} (${latency})\n",
				TimeFormat: "Mon, 2 Jan 2006 15:04:05 MST",
				Output:     os.Stdout,
			},
		),
	)
}
