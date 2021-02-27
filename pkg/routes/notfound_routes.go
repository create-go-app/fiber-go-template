package routes

import (
	"github.com/gofiber/fiber/v2"
)

// NotFoundRoute route for 404 Error.
func NotFoundRoute(a *fiber.App) {
	// Register new route.
	a.Use(
		// Anonimus function.
		func(c *fiber.Ctx) error {
			// Return HTTP 404 status and JSON response.
			if err := c.Status(404).JSON(fiber.Map{
				"error": true,
				"msg":   "sorry, endpoint not found",
			}); err != nil {
				return err
			}

			return nil
		},
	)
}
