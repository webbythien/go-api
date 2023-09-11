package routes

import (
	"github.com/gofiber/fiber/v2"
	"lido-core/v1/app/controllers"
)

func LiveRoute(a *fiber.App) {
	route := a.Group("/live")
	route.Get("/detail", controllers.LiveDetails)
}
