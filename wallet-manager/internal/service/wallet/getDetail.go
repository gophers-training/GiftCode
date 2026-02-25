package walletsvc

import (
	"wallet-manager/internal/model"
	inmemory "wallet-manager/internal/repository/memory"
)

func GetDetails(mobileNo string) (model.Wallet, error) {
	return inmemory.GetWalletDetails(mobileNo)
}
