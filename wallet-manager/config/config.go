package config

import (
	"os"

	"github.com/ghodss/yaml"
)

// AppConfig holds the global configuration for the application.
var AppConfig *Config

// Config represents the application configuration structure.
type Config struct {
	Debug bool      `yaml:"debug"`
	API   APIConfig `yaml:"api"`
}

// LoadConfig reads the configuration from the config.yaml file and populates the AppConfig variable.
func LoadConfig() {
	AppConfig = &Config{}

	configContent, err := os.ReadFile("./config.yaml")
	if err != nil {
		errorMessage := "Failed to read config file: " + err.Error()
		panic(errorMessage)
	}

	err = yaml.Unmarshal(configContent, &AppConfig)
	if err != nil {
		errorMessage := "Failed to parse config file: " + err.Error()
		panic(errorMessage)
	}

	if AppConfig.Debug {
		println("Config loaded successfully:")
		println(string(configContent))
	}

	return

}
