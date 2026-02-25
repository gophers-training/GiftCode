package wallethandler

import (
	"wallet-manager/internal/model"
	walletsvc "wallet-manager/internal/service/wallet"

	"github.com/gofiber/fiber/v3"
)

func List(c fiber.Ctx) error {
	walletList, err := walletsvc.List()
	if err != nil {
		errorMessage := model.FailureResponse{
			Message: "Failed to retrieve wallet list",
			Error:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errorMessage)
	}

	if len(walletList) == 0 {
		response := model.ResponseWithData{
			Message: "No wallets found",
			Data:    []model.Wallet{},
		}
		return c.Status(fiber.StatusOK).JSON(response)
	}

	response := model.ResponseWithData{
		Message: "Wallet list retrieved successfully",
		Data:    walletList,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
