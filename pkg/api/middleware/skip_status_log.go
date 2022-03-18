package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func SkipStatusLog(c *fiber.Ctx) bool {
	return c.Path() == "/status"
}
