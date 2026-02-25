package main

import (
	"sync"
	"wallet-manager/config"
	apisvc "wallet-manager/internal/api"
)

func main() {
	appWG := &sync.WaitGroup{}

	// Load configuration
	config.LoadConfig()

	// Start API server
	appWG.Add(1)
	go apisvc.StartAPIServer(appWG)

	// Add more services here as needed
	// e.g., database connections, background workers, etc.
	// ...

	// Wait for all goroutines to finish
	appWG.Wait()

}
