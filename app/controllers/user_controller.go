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
// @Description Get all exists users.
// @Summary get all exists users
// @Tags Public
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /api/public/users [get]
func GetUsers(c *fiber.Ctx) error {
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all users.
	users, err := db.GetUsers()
	if err != nil {
		// Return, if users not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
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
// @Description Get user by given ID.
// @Summary get user by given ID
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get user by ID.
	user, err := db.GetUser(id)
	if err != nil {
		// Return, if user not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
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
// @Description Create a new user.
// @Summary create a new user
// @Tags Private
// @Accept json
// @Produce json
// @Param email body string true "E-mail"
// @Success 200 {object} models.User
// @Router /api/private/user [post]
func CreateUser(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get data from JWT.
	token := c.Locals("jwt").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	// Set expiration time from JWT data of current user.
	expires := claims["expires"].(int64)

	// Set credential `user:create` from JWT data of current user.
	credential := claims["user:create"].(bool)

	// Create a new user struct.
	user := &models.User{}

	// Checking received data from JSON body.
	if err := c.BodyParser(user); err != nil {
		// Return, if JSON data is not correct.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Only user with `user:create` credential can create a new user profile.
	if credential && now < expires {
		// Create a new validator for user model.
		validate := validators.UserValidator()

		// Validate user fields.
		if err := validate.Struct(user); err != nil {
			// Return, if some fields are not valid.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   utils.ValidatorErrors(err),
			})
		}

		// Create database connection.
		db, err := database.OpenDBConnection()
		if err != nil {
			// Return status 500 and database connection error.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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

		// Create a new user with validated data.
		if err := db.CreateUser(user); err != nil {
			// Return status 500 and create user process error.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
	} else {
		// Return status 403 and permission denied error.
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied, check credentials or expiration time of your token",
			"user":  nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
}

// UpdateUser func for updates user by given ID.
// @Description Update user.
// @Summary update user
// @Tags Private
// @Accept json
// @Produce json
// @Param id body string true "User ID"
// @Success 200 {object} models.User
// @Router /api/private/user [patch]
func UpdateUser(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get data from JWT.
	token := c.Locals("jwt").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	// Set expiration time from JWT data of current user.
	expires := claims["expires"].(int64)

	// Set credential `user:update` from JWT data of current user.
	credential := claims["user:update"].(bool)

	// Create a new user struct.
	user := &models.User{}

	// Checking received data from JSON body.
	if err := c.BodyParser(user); err != nil {
		// Return, if JSON data is not correct.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if user with given ID is exists.
	if _, err := db.GetUser(user.ID); err != nil {
		// Return status 404 and user not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user not found",
		})
	}

	// Only user with `user:update` credential can update user profile.
	if credential && now < expires {
		// Create a new validator for user model.
		validate := validators.UserValidator()

		// Validate user fields.
		if err := validate.Struct(user); err != nil {
			// Return, if some fields are not valid.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   utils.ValidatorErrors(err),
			})
		}

		// Set user data to update:
		user.UpdatedAt = time.Now()

		// Update user.
		if err := db.UpdateUser(user); err != nil {
			// Return status 500 and user update error.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
	} else {
		// Return status 403 and permission denied error.
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied, check credentials or expiration time of your token",
			"user":  nil,
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
}

// DeleteUser func for deletes user by given ID.
// @Description Delete user by given ID.
// @Summary delete user by given ID
// @Tags Private
// @Accept json
// @Produce json
// @Param id body string true "User ID"
// @Success 200 {string} string "ok"
// @Router /api/private/user [delete]
func DeleteUser(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get data from JWT.
	token := c.Locals("jwt").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	// Set expiration time from JWT data of current user.
	expires := claims["expires"].(int64)

	// Set credential `user:delete` from JWT data of current user.
	credential := claims["user:delete"].(bool)

	// Create new User struct
	user := &models.User{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(user); err != nil {
		// Return status 500 and JSON parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if user with given ID is exists.
	if _, err := db.GetUser(user.ID); err != nil {
		// Return status 404 and user not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user not found",
		})
	}

	// Only user with `user:delete` credential can delete user profile.
	if credential && now < expires {
		// Delete user by given ID.
		if err := db.DeleteUser(user.ID); err != nil {
			// Return status 500 and delete user process error.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
	} else {
		// Return status 403 and permission denied error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied, check credentials or expiration time of your token",
		})

	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
	})
}
