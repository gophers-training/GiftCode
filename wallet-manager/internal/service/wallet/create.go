package walletsvc

import (
	"wallet-manager/internal/model"
	inmemory "wallet-manager/internal/repository/memory"
)

func Create(w model.CreateWalletRequest) (model.Wallet, error) {
	newWallet := model.Wallet{
		Mobile: w.Mobile,
	}
	err := inmemory.CreateWallet(newWallet)
	if err != nil {
		return model.Wallet{}, err
	}

	createdWallet, err := inmemory.GetWalletDetails(w.Mobile)
	if err != nil {
		return model.Wallet{}, err
	}

	return createdWallet, nil
}
