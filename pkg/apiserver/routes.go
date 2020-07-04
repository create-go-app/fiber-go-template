package apiserver

import "github.com/gofiber/fiber"

// Routes list of the available routes for project
func Routes(app *fiber.App) {
	// Create group for API routes
	APIGroup := app.Group("/api")

	// API routes
	APIGroup.Get("/docs", func(c *fiber.Ctx) {
		// Set JSON data
		data := fiber.Map{
			"message": "ok",
			"results": []fiber.Map{
				{
					"name": "Documentation",
					"url":  "https://create-go.app/",
				},
				{
					"name": "Detailed guides",
					"url":  "https://create-go.app/detailed-guides/",
				},
				{
					"name": "GitHub",
					"url":  "https://github.com/create-go-app/cli",
				},
			},
		}

		// Set 200 OK status and return JSON, or skip route
		if errJSON := c.Status(200).JSON(data); errJSON != nil {
			// Send 500 status and error to Fiber
			c.Status(500).Next(errJSON)
		}
	})
}
