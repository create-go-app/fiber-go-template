package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateNewJWTAccessToken func for generate a new JWT access (private) token
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

// GenerateNewJWTRefreshToken func for generate a new JWT refresh (public) token.
func GenerateNewJWTRefreshToken() (string, error) {
	// Create a new SHA256 hash.
	sha256 := sha256.New()

	// Create a new now date and time string.
	now := time.Now().String()

	// See: https://pkg.go.dev/io#Writer.Write
	_, err := sha256.Write([]byte(now))
	if err != nil {
		// Return error, it refresh token generation failed.
		return "", err
	}

	return hex.EncodeToString(sha256.Sum(nil)), nil
}
