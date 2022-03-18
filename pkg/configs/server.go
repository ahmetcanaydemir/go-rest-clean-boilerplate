package configs

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	App    *fiber.App
	Config struct {
		Env      string
		Port     string
		Location *time.Location
	}
}

var Server server
