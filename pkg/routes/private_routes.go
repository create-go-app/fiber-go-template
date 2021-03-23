package routes

import (
	"github.com/create-go-app/fiber-go-template/app/controllers"
	"github.com/create-go-app/fiber-go-template/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/private", middleware.JWTProtected())

	// Routes for POST method:
	route.Post("/user", controllers.CreateUser) // create a new user

	// Routes for PUT method:
	route.Put("/user", controllers.UpdateUser) // update user by ID

	// Routes for DELETE method:
	route.Delete("/user", controllers.DeleteUser) // delete one user by ID
}
