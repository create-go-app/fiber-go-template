package controllers

import (
	"github.com/create-go-app/fiber-go-template/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetUsers func gets all exists users or 404 error.
func GetUsers(c *fiber.Ctx) error {
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all users.
	users, err := db.GetUsers()
	if err != nil {
		// Return, if users not found.
		return c.Status(404).JSON(fiber.Map{
			"error": true,
			"msg":   "users were not found",
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

// GetUser func gets one user by given ID or 404 error.
func GetUser(c *fiber.Ctx) error {
	// Catch user ID from URL.
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get user by ID.
	user, err := db.GetUser(id)
	if err != nil {
		// Return, if user not found.
		return c.Status(404).JSON(fiber.Map{
			"error": true,
			"msg":   "user with the given ID is not found",
			"user":  nil,
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
}
