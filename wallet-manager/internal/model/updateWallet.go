package model

type UpdateWalletRequest struct {
	Amount    float64 `json:"amount" validate:"required"`
	Operation string  `json:"operation" validate:"required,oneof=add subtract set"`
}
