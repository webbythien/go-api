package controllers

import (
	"lido-core/v1/app/schemas"
	"lido-core/v1/app/services"

	"github.com/gofiber/fiber/v2"
)

func LiveDetails(c *fiber.Ctx) error {
	typ := c.Query("type")
	switch typ {
	case "main":
		live := services.LiveMain()
		return c.JSON(schemas.DataResponse{
			Data:    live,
			Success: true,
			Msg:     "ok",
		})
	case "recommend":
		recommend, err := services.RecommendVideo(9)
		if err != nil {
			return c.JSON(schemas.DataResponse{
				Data:    []interface{}{},
				Success: true,
				Msg:     "ok",
			})
		}
		return c.JSON(schemas.DataResponse{
			Data:    recommend,
			Success: true,
			Msg:     "ok",
		})
	}
	return c.Status(fiber.StatusBadRequest).JSON(schemas.DataResponse{
		Msg:     "bad request",
		Success: false,
		Data:    nil,
	})
}
