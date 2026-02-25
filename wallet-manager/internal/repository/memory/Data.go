package inmemory

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
	"wallet-manager/config"
	"wallet-manager/internal/model"
)

var Data *InMemoryData

func Init() {
	Data = &InMemoryData{
		Wallets:    map[string]model.Wallet{},
		WalletMutx: sync.Mutex{},
	}

	CreateInitWallets()

	if config.AppConfig.Debug {
		jsonContent, _ := json.MarshalIndent(Data.Wallets, "", "  ")
		fmt.Printf("Loaded Data:\n %s", string(jsonContent))
	}
}

func CreateInitWallets() {
	wallets := []model.Wallet{
		{
			Mobile:    "09128888888",
			Balance:   100000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Mobile:    "09126666666",
			Balance:   200000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, w := range wallets {
		CreateWallet(w)
	}
}
