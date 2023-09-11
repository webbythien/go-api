package controllers

import (
	"github.com/gofiber/fiber/v2"
	"lido-core/v1/app/services"
)

// HealthCheck health check api server.
// @Description health check api server.
// @Tags HealthCheck
// @Accept json
// @Produce json
// @Success 200
// @Router /healthcheck [GET]
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "OK",
	})
}

func HealthCheckWorker(c *fiber.Ctx) error {
	err := services.HealthCheck()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "OK",
	})
}
