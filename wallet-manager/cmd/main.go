package main

import (
	"sync"
	"wallet-manager/config"
	apisvc "wallet-manager/internal/api"
	inmemory "wallet-manager/internal/repository/memory"
)

func main() {
	appWG := &sync.WaitGroup{}

	// Load configuration
	config.LoadConfig()

	// Load init wallets data
	inmemory.Init()

	// Start API server
	appWG.Add(1)
	go apisvc.StartAPIServer(appWG)

	// Add more services here as needed
	// e.g., database connections, background workers, etc.
	// ...

	// Wait for all goroutines to finish
	appWG.Wait()

}
