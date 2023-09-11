package routes

import (
	"lido-core/v1/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func WalletRoute(a *fiber.App) {
	route := a.Group("/wallet")
	route.Get("/balance/:address", controllers.BalanceController)
}
