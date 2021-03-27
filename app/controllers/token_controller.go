package controllers

import (
	"context"
	"time"

	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/create-go-app/fiber-go-template/pkg/utils"
	"github.com/create-go-app/fiber-go-template/platform/cache"
	"github.com/create-go-app/fiber-go-template/platform/database"

	"github.com/gofiber/fiber/v2"
)

// RenewTokens method for renew access and refresh tokens.
// @Description Renew access and refresh tokens.
// @Summary renew access and refresh tokens
// @Tags Token
// @Accept json
// @Produce json
// @Param refresh_token body string true "Refresh token"
// @Success 200 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/token/renew [post]
func RenewTokens(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current user.
	expiresAccessToken := claims.Expires

	// Checking, if now time greather than Access token expiration time.
	if now > expiresAccessToken {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create a new renew refresh token struct.
	renew := &models.Renew{}

	// Checking received data from JSON body.
	if err := c.BodyParser(renew); err != nil {
		// Return, if JSON data is not correct.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from Refresh token of current user.
	expiresRefreshToken, err := utils.ParseRefreshToken(renew.RefreshToken)
	if err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if now time greather than Refresh token expiration time.
	if now < expiresRefreshToken {
		// Define user ID.
		userID := claims.UserID

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
		foundedUser, err := db.GetUserByID(userID)
		if err != nil {
			// Return, if user not found.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "user with the given ID is not found",
			})
		}

		// Get role credentials from founded user.
		credentials, err := utils.GetCredentialsByRole(foundedUser.UserRole)
		if err != nil {
			// Return status 400 and error message.
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		// Generate JWT Access & Refresh tokens.
		tokens, err := utils.GenerateNewTokens(userID.String(), credentials)
		if err != nil {
			// Return status 500 and token generation error.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		// Create a new Redis connection.
		connRedis, err := cache.RedisConnection()
		if err != nil {
			// Return status 500 and Redis connection error.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		// Save refresh token to Redis.
		errRedis := connRedis.Set(context.Background(), userID.String(), tokens.Refresh, 0).Err()
		if errRedis != nil {
			// Return status 500 and Redis connection error.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   errRedis.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"error": false,
			"msg":   nil,
			"tokens": fiber.Map{
				"access":  tokens.Access,
				"refresh": tokens.Refresh,
			},
		})
	} else {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, your session was ended earlier",
		})
	}
}
