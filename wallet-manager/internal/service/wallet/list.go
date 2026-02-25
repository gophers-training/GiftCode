package walletsvc

import (
	"wallet-manager/internal/model"
	inmemory "wallet-manager/internal/repository/memory"
)

func List() ([]model.Wallet, error) {
	return inmemory.ListWallets()
}
