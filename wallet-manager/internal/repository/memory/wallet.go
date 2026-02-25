package inmemory

import (
	"errors"
	"fmt"
	"wallet-manager/internal/model"
)

func CreateWallet(w model.Wallet) error {
	Data.WalletMutx.Lock()
	defer Data.WalletMutx.Unlock()

	Data.Wallets[w.Mobile] = w

	return nil
}

func DeleteWallet(mobileNo string) error {
	Data.WalletMutx.Lock()
	defer Data.WalletMutx.Unlock()

	delete(Data.Wallets, mobileNo)

	return nil
}

func GetWalletDetails(mobileNo string) (model.Wallet, error) {
	Data.WalletMutx.Lock()
	defer Data.WalletMutx.Unlock()

	wallet, exists := Data.Wallets[mobileNo]
	if !exists {
		errorMessage := "Wallet with mobile number " + mobileNo + " not found"
		return model.Wallet{}, errors.New(errorMessage)
	}

	return wallet, nil
}

func ListWallets() ([]model.Wallet, error) {
	Data.WalletMutx.Lock()
	defer Data.WalletMutx.Unlock()

	wallets := make([]model.Wallet, 0, len(Data.Wallets))
	for _, wallet := range Data.Wallets {
		wallets = append(wallets, wallet)
	}

	return wallets, nil
}

func UpdateWallet(mobileNo string, amount float64, operation string) (model.Wallet, error) {
	Data.WalletMutx.Lock()
	defer Data.WalletMutx.Unlock()

	wallet, exists := Data.Wallets[mobileNo]
	if !exists {
		errorMessage := "Wallet with mobile number " + mobileNo + " not found"
		return model.Wallet{}, errors.New(errorMessage)
	}

	switch operation {
	case "add":
		wallet.Balance += amount
	case "subtract":
		wallet.Balance -= amount
	default:
		errorMessage := fmt.Sprintf("invalid operation %s", operation)
		return model.Wallet{}, errors.New(errorMessage)
	}

	Data.Wallets[mobileNo] = wallet

	return wallet, nil
}
