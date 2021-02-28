package controllers

import (
	"github.com/create-go-app/fiber-go-template/platform/database"
	"github.com/gofiber/fiber/v2"
)

// GetUsers func gets all exists users or 404 error.
func GetUsers(c *fiber.Ctx) error {
	// Create DB connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return database connection error.
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all users.
	users, err := db.GetUsers()
	if err != nil {
		// Return, if Users not found.
		return c.Status(404).JSON(fiber.Map{
			"error": true,
			"msg":   "users not found",
			"count": 0,
			"users": nil,
		})

	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(users),
		"users": users,
	})
}
