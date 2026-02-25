package wallethandler

import (
	"wallet-manager/internal/model"
	walletsvc "wallet-manager/internal/service/wallet"

	"github.com/gofiber/fiber/v3"
)

func Create(c fiber.Ctx) error {
	createRequest := model.CreateWalletRequest{}
	if err := c.Bind().JSON(&createRequest); err != nil {
		errorMessage := model.FailureResponse{
			Message: "Invalid request payload",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(errorMessage)
	}

	createWallet, err := walletsvc.Create(createRequest)
	if err != nil {
		errorMessage := model.FailureResponse{
			Message: "Failed to create wallet",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errorMessage)
	}

	response := model.ResponseWithData{
		Message: "Wallet created successfully",
		Data:    createWallet,
	}
	c.Status(fiber.StatusCreated).JSON(response)

	return nil
}
