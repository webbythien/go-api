package controllers

import (
	"errors"
	"lido-core/v1/app/schemas"
	"lido-core/v1/app/services"
	"lido-core/v1/pkg/utils"
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber/v2"
)

func CreateQuiz(c *fiber.Ctx) error {
	request := new(schemas.Quiz)
	if err := c.BodyParser(request); err != nil {
		log.Printf("Error parsing request body: " + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Msg:     "error parsing request body",
			Success: false,
		})
	}
	if err := services.CreateQuiz(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Msg:     err.Error(),
			Success: false,
		})
	}

	return c.JSON(schemas.Response{
		Msg:     "ok",
		Success: true,
	})
}

func GetQuiz(c *fiber.Ctx) error {
	id := c.Params("lid")
	res, err := services.GetQuiz(id)
	if err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return c.Status(fiber.StatusOK).JSON(schemas.DataResponse{
				Msg:     "no questions open",
				Success: true,
				Data:    nil,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.DataResponse{
			Msg:     err.Error(),
			Success: false,
			Data:    nil,
		})
	}
	return c.JSON(schemas.DataResponse{
		Msg:     "ok",
		Success: true,
		Data:    res,
	})
}

func QuizAnswer(c *fiber.Ctx) error {
	request := new(schemas.AnswerRequest)
	if err := c.BodyParser(request); err != nil {
		log.Printf("Error parsing request body: " + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Msg:     "error parsing request body",
			Success: false,
		})
	}
	address, err := utils.GetUserAddress(c)
	if err != nil {
		log.Fatal(err)
	}
	err = services.QuizAnswer(request.Qid, request.Aid, address)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Msg:     err.Error(),
			Success: false,
		})
	}
	return c.JSON(schemas.Response{
		Msg:     "ok",
		Success: true,
	})
}

func QuizStart(c *fiber.Ctx) error {
	id := c.Params("qid")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Msg:     "missing field: qid",
			Success: false,
		})
	}
	err := services.QuizStart(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.Response{
			Msg:     err.Error(),
			Success: false,
		})
	}
	return c.JSON(schemas.Response{
		Msg:     "ok",
		Success: true,
	})
}

func QuizClose(c *fiber.Ctx) error {
	id := c.Params("qid")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Msg:     "missing field: qid",
			Success: false,
		})
	}
	err := services.QuizClose(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Msg:     err.Error(),
			Success: false,
		})
	}
	return c.JSON(schemas.Response{
		Msg:     "ok",
		Success: true,
	})
}

func QuizStat(c *fiber.Ctx) error {
	id := c.Params("qid")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.DataResponse{
			Msg:     "missing field: qid",
			Success: false,
			Data:    nil,
		})
	}
	address := utils.GetUserAddressOtps(c)
	res, err := services.QuizStat(id, address)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.DataResponse{
			Msg:     err.Error(),
			Success: false,
			Data:    nil,
		})
	}
	return c.JSON(schemas.DataResponse{
		Msg:     "ok",
		Success: true,
		Data:    res,
	})
}

func QuizActive(c *fiber.Ctx) error {
	id := c.Params("lid")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.Response{
			Msg:     "missing field: lid",
			Success: false,
		})
	}
	err := services.QuizActivate(id)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(schemas.Response{
			Msg:     err.Error(),
			Success: true,
		})
	}
	return c.JSON(schemas.Response{
		Msg:     "ok",
		Success: true,
	})
}
