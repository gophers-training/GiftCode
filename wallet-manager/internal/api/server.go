package apisvc

import (
	"sync"
	"wallet-manager/config"

	"github.com/gofiber/fiber/v2"
)

func StartAPIServer(wg *sync.WaitGroup) {
	defer wg.Done()

	api := fiber.New()

	err := api.Listen(config.AppConfig.API.GetListenAddress())
	if err != nil {
		panic("Failed to start API server: " + err.Error())
	}
}
