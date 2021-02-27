package routes

import "github.com/gofiber/fiber/v2"

// PublicRoutes public routes group for all users.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/public")

	// Routes for GET method.
	route.Get("/users", nil)
}
