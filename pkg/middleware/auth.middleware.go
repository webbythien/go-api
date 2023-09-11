package middleware

import (
	"lido-core/v1/app/schemas"
	"lido-core/v1/pkg/constants"
	"lido-core/v1/pkg/utils"
	"log"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func SignInValidate(c *fiber.Ctx) error {
	validate := validator.New()
	request := new(schemas.SignIn)
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
	c.Locals(constants.SignInData, *request)
	return c.Next()
}

func Authenticate(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"msg":     "unauthorized",
			"success": false,
		})
	}
	_, err := utils.ParseToken(authHeader)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"msg":     "unauthorized",
			"success": false,
		})
	}
	return c.Next()
}
