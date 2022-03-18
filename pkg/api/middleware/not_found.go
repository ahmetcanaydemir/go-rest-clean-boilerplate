package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func NotFound(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	return c.Status(fiber.StatusNotFound).SendString("404: Page not Found")
}
