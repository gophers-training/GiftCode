package walletsvc

import (
	"wallet-manager/internal/model"
	inmemory "wallet-manager/internal/repository/memory"
)

func Create(w model.Wallet) error {
	return inmemory.CreateWallet(w)
}
