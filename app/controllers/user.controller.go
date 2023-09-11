package controllers

import (
	"lido-core/v1/app/schemas"
	"lido-core/v1/app/services"
	"lido-core/v1/pkg/constants"
	"lido-core/v1/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func SignIn(c *fiber.Ctx) error {
	request, ok := c.Locals(constants.SignInData).(schemas.SignIn)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.UserSignInResponse{
			AccessToken: "",
			Message:     "Failed to load data",
			Success:     false,
		})
	}
	token, err := services.SignIn(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.UserSignInResponse{
			AccessToken: "",
			Message:     err.Error(),
			Success:     false,
		})
	}
	return c.Status(fiber.StatusOK).JSON(schemas.UserSignInResponse{
		AccessToken: token,
		Message:     "ok",
		Success:     true,
	})
}

func GetMessage(c *fiber.Ctx) error {
	msg := services.GenerateMessage()
	return c.Status(fiber.StatusOK).JSON(schemas.DataResponse{
		Msg:     "ok",
		Success: true,
		Data:    msg,
	})
}

func GetQuizResult(c *fiber.Ctx) error {
	id := c.Params("lid")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.DataResponse{
			Msg:     "missing field: live id",
			Success: false,
			Data:    nil,
		})
	}
	address, err := utils.GetUserAddress(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.DataResponse{
			Msg:     err.Error(),
			Success: false,
			Data:    nil,
		})
	}
	data, err := services.GetQuizResult(address, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.DataResponse{
			Msg:     err.Error(),
			Success: false,
			Data:    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(schemas.DataResponse{
		Msg:     "ok",
		Success: true,
		Data:    data,
	})
}
