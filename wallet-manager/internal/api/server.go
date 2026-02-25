package apisvc

import (
	"strings"
	"sync"
	"wallet-manager/config"
	wallethandler "wallet-manager/internal/api/handler/wallet"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func StartAPIServer(wg *sync.WaitGroup) {
	defer wg.Done()

	api := fiber.New()

	// Set CORS configuration
	SetCorsConfig(api)

	// Setup Routes
	SetAPIRoutes(api)

	err := api.Listen(config.AppConfig.API.GetListenAddress())
	if err != nil {
		panic("Failed to start API server: " + err.Error())
	}
}

func SetCorsConfig(api *fiber.App) {
	api.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(config.AppConfig.API.CORS.AllowedOrigins, ","),
		AllowMethods:     strings.Split(config.AppConfig.API.CORS.AllowedMethods, ","),
		AllowHeaders:     strings.Split(config.AppConfig.API.CORS.AllowedHeaders, ","),
		AllowCredentials: config.AppConfig.API.CORS.AllowCredentials,
	}))
}

func SetAPIRoutes(api *fiber.App) {
	// Get Wallet List
	api.Get("/wallet", wallethandler.List)
	// Get Wallet Detail
	api.Get("/wallet/:id", wallethandler.GetDetail)
	// Create Wallet
	api.Post("/wallet", wallethandler.Create)
	// Update Wallet
	api.Put("/wallet/:id", wallethandler.Update)
	// Delete Wallet
	api.Delete("/wallet/:id", wallethandler.Delete)
}
