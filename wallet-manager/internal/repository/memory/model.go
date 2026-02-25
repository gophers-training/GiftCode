package inmemory

import (
	"sync"
	"wallet-manager/internal/model"
)

type InMemoryData struct {
	Wallets    map[string]model.Wallet
	WalletMutx sync.Mutex
}
