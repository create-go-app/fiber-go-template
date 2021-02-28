package routes

import (
	"github.com/create-go-app/fiber-go-template/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes public routes group for all users.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/public")

	// Routes for GET method.
	route.Get("/users", controllers.GetUsers) // get list of all users
}
