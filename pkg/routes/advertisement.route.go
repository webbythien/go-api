package routes

import (
	"github.com/gofiber/fiber/v2"
	"lido-core/v1/app/controllers"
)

func AdvertisementRoute(a *fiber.App) {
	route := a.Group("/advertisement")
	route.Get("/all", controllers.AllAdvertisement)
}
