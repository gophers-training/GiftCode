package model

type CreateWalletRequest struct {
	Mobile string `json:"mobile" validate:"required,e164"`
}
