package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerHost string
	ServerPort string
	Env        string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	config := &Config{
		ServerHost: getEnv("SERVER_HOST", "localhost"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
		Env:        getEnv("ENV", "development"),
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
