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
func GenerateNewJWTAccessToken(credentials []string, id string) (string, error) {
	// Catch JWT secret key from .env file.
	secret := os.Getenv("JWT_SECRET_KEY")

	// Create a new JWT access token and claims.
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// Set public claims:
	claims["id"] = id
	claims["expires"] = time.Now().Add(time.Hour * 72).Unix()

	// Set private token credentials:
	for _, credential := range credentials {
		claims[credential] = true
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

	// Create a new now date and time string with salt.
	refresh := os.Getenv("JWT_REFRESH_KEY") + time.Now().String()

	// See: https://pkg.go.dev/io#Writer.Write
	_, err := sha256.Write([]byte(refresh))
	if err != nil {
		// Return error, it refresh token generation failed.
		return "", err
	}

	return hex.EncodeToString(sha256.Sum(nil)), nil
}
