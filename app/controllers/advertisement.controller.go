package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"lido-core/v1/app/schemas"
	"lido-core/v1/app/services"
)

func AllAdvertisement(c *fiber.Ctx) error {
	fmt.Println("Hello, World!")
	allAdvertisements, err := services.GetAllAdvertisements()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.DataResponse{
			Data:    nil,
			Success: false,
			Msg:     err.Error(), // Include the error message in the response
		})
	}
	return c.JSON(schemas.DataResponse{
		Data:    allAdvertisements,
		Success: true,
		Msg:     "ok",
	})
}
