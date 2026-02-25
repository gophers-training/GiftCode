package config

import "fmt"

// APIConfig holds the configuration for the API server.
type APIConfig struct {
	ListenAddress string        `yaml:"listenAddress"`
	ListenPort    int           `yaml:"listenPort"`
	CORS          APICORSConfig `yaml:"cors"`
}

// GetListenAddress returns the full listen address in the format "address:port".
func (a *APIConfig) GetListenAddress() string {
	return fmt.Sprintf("%s:%d", a.ListenAddress, a.ListenPort)
}

// APICORSConfig holds the CORS configuration for the API server.
type APICORSConfig struct {
	AllowedOrigins   string `yaml:"allowedOrigins"`
	AllowedMethods   string `yaml:"allowedMethods"`
	AllowedHeaders   string `yaml:"allowedHeaders"`
	AllowCredentials bool   `yaml:"allowCredentials"`
}
