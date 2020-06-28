package apiserver

import "github.com/gofiber/fiber"

// Routes ...
func Routes(app *fiber.App) {
	// Create group for API routes
	APIGroup := app.Group("/api")

	// API routes
	APIGroup.Get("/posts", func(c *fiber.Ctx) {
		// Set JSON data
		data := fiber.Map{
			"message": nil,
			"results": nil,
		}

		// Set 200 OK status and return JSON, or skip route
		if errJSON := c.Status(200).JSON(data); errJSON != nil {
			// Send 500 status and error to Fiber
			c.Status(500).Next(errJSON)
		}
	})
}
