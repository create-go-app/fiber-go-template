package controllers

import (
	"time"

	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/create-go-app/fiber-go-template/app/validators"
	"github.com/create-go-app/fiber-go-template/pkg/utils"
	"github.com/create-go-app/fiber-go-template/platform/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetUsers func gets all exists users.
// @Description Gets all exists users.
// @Tags Public
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /api/public/users [get]
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
// @Description Gets one user by given ID.
// @Tags Public
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /api/public/user/{id} [get]
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

// CreateUser func for creates a new user.
func CreateUser(c *fiber.Ctx) error {
	// Create a new user struct.
	user := &models.User{}

	// Checking received data from JSON body.
	if err := c.BodyParser(user); err != nil {
		// Return, if JSON data is not correct.
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set initialized default data for user:
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Time{}
	user.UserStatus = 1 // 0 == blocked, 1 == active
	user.UserAttrs = models.UserAttrs{}

	// Create a new validator for user model.
	validate := validators.UserValidator()

	// Validate user fields.
	if err := validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
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

	// Create a new user with validated data.
	if err := db.CreateUser(user); err != nil {
		// Return status 500 and create user process error.
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
}

// UpdateUser func for updates user by given ID.
func UpdateUser(c *fiber.Ctx) error {
	// Get data from JWT.
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	// Set admin status from JWT data of current user.
	isAdmin := claims["is_admin"].(bool)

	// Create a new user struct.
	user := &models.User{}

	// Checking received data from JSON body.
	if err := c.BodyParser(user); err != nil {
		// Return, if JSON data is not correct.
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

	// Checking, if user with given ID is exists.
	if _, err := db.GetUser(user.ID); err != nil {
		// Return status 404 and user not found error.
		return c.Status(404).JSON(fiber.Map{
			"error": true,
			"msg":   "user not found",
		})
	}

	// Get ID from JWT of current user and convert it to UUID.
	currentUserID, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		// Return status 500 and UUID parcing error.
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Only user itself or admin can update user profile.
	if currentUserID == user.ID || isAdmin {
		// Create a new validator for user model.
		validate := validators.UserValidator()

		// Validate user fields.
		if err := validate.Struct(user); err != nil {
			// Return, if some fields are not valid.
			return c.Status(500).JSON(fiber.Map{
				"error": true,
				"msg":   utils.ValidatorErrors(err),
			})
		}

		// Set user data to update:
		user.UpdatedAt = time.Now()

		// Update user.
		if err := db.UpdateUser(user); err != nil {
			// Return status 500 and user update error.
			return c.Status(500).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
	} else {
		// Return status 403 and permission denied error.
		return c.Status(403).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied",
			"user":  nil,
		})
	}

	return c.Status(202).JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
}

// DeleteUser func deletes user by given ID.
func DeleteUser(c *fiber.Ctx) error {
	// Catch data from JWT.
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	// Set admin status.
	isAdmin := claims["is_admin"].(bool)

	// Create new User struct
	user := &models.User{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(user); err != nil {
		// Return status 500 and JSON parse error.
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

	// Checking, if user with given ID is exists.
	if _, err := db.GetUser(user.ID); err != nil {
		// Return status 404 and user not found error.
		return c.Status(404).JSON(fiber.Map{
			"error": true,
			"msg":   "user not found",
		})
	}

	// Check, if current user request from admin.
	if isAdmin {
		// Delete user by given ID.
		if err := db.DeleteUser(user.ID); err != nil {
			// Return status 500 and delete user process error.
			return c.Status(500).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
	} else {
		// Return status 403 and permission denied error.
		return c.Status(403).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied",
		})

	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
	})
}
