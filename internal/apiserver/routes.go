package apiserver

import "github.com/gofiber/fiber"

// IndexHandler ...
func IndexHandler(c *fiber.Ctx) {
	c.Send("Hello, World!")
}
