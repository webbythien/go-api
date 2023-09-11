package routes

import (
	"lido-core/v1/app/controllers"
	"lido-core/v1/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(a *fiber.App) {
	// Query Params:
	// lid: live ID (id youtube)
	route := a.Group("/user")
	route.Post("/login", middleware.SignInValidate, controllers.SignIn)
	route.Get("/message", controllers.GetMessage)
	route.Get("/quiz/:lid", middleware.Authenticate, controllers.GetQuizResult)
}
