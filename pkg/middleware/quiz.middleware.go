package middleware

import (
	"lido-core/v1/app/schemas"
	"log"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func QuizValidator(c *fiber.Ctx) error {
	validate := validator.New()
	request := new(schemas.Quiz)
	if err := c.BodyParser(request); err != nil {
		log.Printf("Error parsing request body: " + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":     "error parsing request body",
			"success": false,
		})
	}
	if err := validate.Struct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			//"data":    request,
			"msg":     "missing required fields",
			"success": false,
		})
	}
	return c.Next()
}

func AnswerValidator(c *fiber.Ctx) error {
	validate := validator.New()
	request := new(schemas.AnswerRequest)
	if err := c.BodyParser(request); err != nil {
		log.Printf("Error parsing request body: " + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":     "error parsing request body",
			"success": false,
		})
	}
	if err := validate.Struct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":     "missing required fields",
			"success": false,
		})
	}
	return c.Next()
}
