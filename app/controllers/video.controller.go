package controllers

import (
	"lido-core/v1/app/schemas"
	"lido-core/v1/app/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LandingPage(c *fiber.Ctx) error {
	typ := c.Query("type")
	status := c.Query("status")
	switch typ {
	case "current":
		if status == "off" {
			return c.JSON(schemas.DataResponse{
				Msg:     "ok",
				Success: true,
				Data: schemas.VideoMain{
					Started: time.Now().UTC().Unix(),
					Status:  "off",
				},
			})

		}
		live := services.LiveMain()
		return c.JSON(schemas.DataResponse{
			Msg:     "ok",
			Success: true,
			Data:    live,
		})
	case "previous":
		prev, err := services.RecommendVideo(3)
		if err != nil {
			return c.JSON(schemas.DataResponse{
				Msg:     "ok",
				Success: true,
				Data:    []interface{}{},
			})
		}
		return c.JSON(schemas.DataResponse{
			Msg:     "ok",
			Success: true,
			Data:    prev,
		})
	}

	return c.Status(fiber.StatusBadRequest).JSON(schemas.DataResponse{
		Msg:     "bad request",
		Success: false,
		Data:    nil,
	})

}

func PreviousStream(c *fiber.Ctx) error {
	id := c.Params("id")
	recommend := c.Query("recommend")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.DataResponse{
			Msg:     "missing field: id",
			Success: false,
			Data:    nil,
		})
	}
	if recommend == "true" {
		data, err := services.PreviousStreamRecommend(id)
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
			Data:    data,
		})
	}
	prevStream := services.PreviousStreamDetails(id)
	return c.Status(fiber.StatusOK).JSON(schemas.DataResponse{
		Msg:     "ok",
		Success: true,
		Data:    *prevStream,
	})
}
