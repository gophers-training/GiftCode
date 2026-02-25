package apisvc

import (
	"strings"
	"sync"
	"wallet-manager/config"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func StartAPIServer(wg *sync.WaitGroup) {
	defer wg.Done()

	api := fiber.New()

	// Set CORS configuration
	SetCorsConfig(api)

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
