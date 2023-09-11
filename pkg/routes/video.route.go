package routes

import (
	"lido-core/v1/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func VideoRoute(a *fiber.App) {
	// route := a.Group("/mock")
	a.Get("/landing", controllers.LandingPage)
	a.Get("/previous/:id", controllers.PreviousStream)
}
