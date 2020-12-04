package apiserver

import "github.com/gofiber/fiber/v2"

// Routes list of the available routes for project
func Routes(app *fiber.App) {
	// Create group for API routes
	APIGroup := app.Group("/api")

	// API routes
	APIGroup.Get("/docs", func(c *fiber.Ctx) error {
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

		// Set 200 OK status and return JSON
		return c.Status(200).JSON(data)
	})
}
