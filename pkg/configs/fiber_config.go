package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Checking .env variable, it must be int.
	readTimeout, err := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	if err != nil {
		panic(err)
	}

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeout),
	}
}
