package utils

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

// GenerateNewJWTAccessToken func for generate a new JWT access token
// with user ID and permissions.
func GenerateNewJWTAccessToken(permission, id string) (string, error) {
	// Catch secret JWT value from .env file.
	secret := os.Getenv("JWT_SECRET_TOKEN")

	// Create a new JWT access token and claims.
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// Set public claims:
	claims["id"] = id
	claims["is_admin"] = false

	// Set private claims:
	switch permission {
	case "admin":
		claims["is_admin"] = true
	}

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}
