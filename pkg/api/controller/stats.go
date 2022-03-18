package controller

import (
	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/api/service"
	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/configs"
	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/utils/helper"
	"github.com/gofiber/fiber/v2"
)

type StatsController struct {
	service service.StatsService
}

func NewStatsController() {
	ct := &StatsController{
		service: service.NewStatsService(),
	}
	configs.Server.App.Get("/Stats", ct.GetStats)
}

// GetStats godoc
// @Summary      Get Stats
// @Tags         Stats
// @Description  This endpoint returns the status of the Go Rest Clean Boilerplate
// @Success      200  {object}  model.Stats  "Stats"
// @Failure      404
// @Failure      500
// @Router       /Stats [get]
func (ct StatsController) GetStats(c *fiber.Ctx) error {
	stats := ct.service.GetStats()
	if stats == nil {
		return c.Next()
	}
	return helper.JSON(c, stats)
}
