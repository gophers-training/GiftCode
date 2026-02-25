package walletsvc

import (
	"wallet-manager/internal/model"
	inmemory "wallet-manager/internal/repository/memory"
)

func Update(mobileNo string, amount float64, operation string) (model.Wallet, error) {
	return inmemory.UpdateWallet(mobileNo, amount, operation)
}
