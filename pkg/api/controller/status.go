package controller

import (
	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/configs"
	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/utils/helper"
	"github.com/gofiber/fiber/v2"
)

type StatusController struct {
}

func NewStatusController() {
	ct := &StatusController{}
	configs.Server.App.Get("/status", ct.GetStatus)
}

type Status struct {
	Ping Ping
}
type Ping struct {
	Message   *string
	Error     *error
	IsHealthy bool
}

// GetStatus godoc
// @Summary      Get Status
// @Tags         Status
// @Description  This endpoint returns the status of the Go Rest Clean Boilerplate
// @Success      200
// @Failure      500
// @Router       /status [get]
func (ct StatusController) GetStatus(c *fiber.Ctx) error {
	status := Status{Ping: Ping{Message: nil, Error: nil, IsHealthy: true}}
	return helper.JSON(c, status)
}
