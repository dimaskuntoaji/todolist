package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() error {
	return godotenv.Load()
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}