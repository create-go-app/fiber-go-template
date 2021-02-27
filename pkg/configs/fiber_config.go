package configs

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// FiberConfig ...
func FiberConfig() fiber.Config {
	return fiber.Config{
		ReadTimeout: 60 * time.Second,
	}
}
