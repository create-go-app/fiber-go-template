package routes

import (
	"github.com/create-go-app/fiber-go-template/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/public")

	// Routes for GET method:
	route.Get("/user/:id", controllers.GetUser) // get one user by ID
	route.Get("/users", controllers.GetUsers)   // get list of all users
}
