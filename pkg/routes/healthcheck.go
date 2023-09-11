package routes

import (
	"github.com/gofiber/fiber/v2"
	"lido-core/v1/app/controllers"
)

func HealthCheck(a *fiber.App) {
	a.Get("/healthcheck", controllers.HealthCheck)
	a.Get("/workercheck", controllers.HealthCheckWorker)
}
