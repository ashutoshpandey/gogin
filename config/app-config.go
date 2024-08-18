package config

import (
	"log"

	"github.com/joho/godotenv"
)

// Config struct to hold all configuration variables
type AppConfig struct {
	JWT_SECRET_KEY string
}

// LoadConfig loads environment variables from the .env file
func LoadAppConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config := &AppConfig{
		JWT_SECRET_KEY: getEnv("JWT_SECRET_KEY", "8080"),
	}

	return config
}
