package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"

	jwtAuthMiddleware "github.com/gofiber/jwt/v2"
)

// JWTAuthProtected func for specify routes group with JWT authentication.
func JWTAuthProtected() fiber.Handler {
	// Create config for JWT authentication middleware.
	config := jwtAuthMiddleware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_TOKEN")),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Return status 403 and failed authentication error.
			return c.Status(403).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		},
	}

	return jwtAuthMiddleware.New(config)
}
