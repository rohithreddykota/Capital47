package config // config.go

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Nessie NessieConfig `mapstructure:"nessie"`
	Server ServerConfig `mapstructure:"server"`
}

type NessieConfig struct {
	APIKey  string `mapstructure:"api_key"`
	BaseURL string `mapstructure:"base_url"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

// LoadConfig loads configuration from the given file path
func LoadConfig(filePath string) (*Config, error) {
	viper.SetConfigFile(filePath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %s", err)
	}

	return &config, nil
}

// PrintConfig prints the loaded configuration to the console
func PrintConfig(config *Config) {
	fmt.Printf("Nessie API Key: %s\n", config.Nessie.APIKey)
	fmt.Printf("Nessie Base URL: %s\n", config.Nessie.BaseURL)
	fmt.Printf("Server Port: %d\n", config.Server.Port)
}
