package routes

import (
	"lido-core/v1/app/controllers"
	"lido-core/v1/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func QuizRoute(a *fiber.App) {
	// Query params:
	// qid: question id documentation
	// aid: answer id documentation
	// lid: Live ID (id youtube)
	route := a.Group("/quiz")
	route.Post("/create", middleware.QuizValidator, controllers.CreateQuiz)
	route.Post("/answer", middleware.Authenticate, middleware.AnswerValidator, controllers.QuizAnswer)
	route.Post("/start/:qid", controllers.QuizStart)
	route.Post("/close/:qid", controllers.QuizClose)
	route.Post("/active/:lid", controllers.QuizActive)
	route.Get("/get/:lid", controllers.GetQuiz)
	route.Get("/stat/:qid", controllers.QuizStat)
}
