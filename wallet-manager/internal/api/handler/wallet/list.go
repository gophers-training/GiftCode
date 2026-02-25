package wallethandler

import (
	"wallet-manager/internal/model"
	walletsvc "wallet-manager/internal/service/wallet"

	"github.com/gofiber/fiber/v3"
)

// @Summary Get list of all wallets
// @Description Retrieve a list of all wallets
// @Tags Wallet
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseWithData{data=[]model.Wallet} "Wallet list retrieved successfully"
// @Failure 500 {object} model.FailureResponse "Failed to retrieve wallet list"
// @Router /wallet [get]
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
