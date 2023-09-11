package controllers

import (
	"lido-core/v1/app/schemas"
	"lido-core/v1/app/services"

	"github.com/gofiber/fiber/v2"
)

func BalanceController(c *fiber.Ctx) error {
	address := c.Params("address")
	if address == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.DataResponse{
			Msg:     "Address is required",
			Success: false,
			Data:    nil,
		})
	}
	balance, err := services.BalanceOf(address)
	if err != nil {
		return c.JSON(schemas.DataResponse{
			Data: schemas.BalanceResponse{
				Address: address,
				Balance: 0,
			},
			Msg:     "ok",
			Success: true,
		})
	}
	return c.JSON(schemas.DataResponse{
		Data: schemas.BalanceResponse{
			Address: address,
			Balance: balance,
		},
		Msg:     "ok",
		Success: true,
	})
}
