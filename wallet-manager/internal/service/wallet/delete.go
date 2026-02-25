package walletsvc

import inmemory "wallet-manager/internal/repository/memory"

func Delete(mobileNo string) error {
	return inmemory.DeleteWallet(mobileNo)
}
