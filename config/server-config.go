package config

import (
	"log"

	"github.com/joho/godotenv"
)

// Config struct to hold all configuration variables
type ServerConfig struct {
	PORT            string
	ALLOWED_ORIGINS string
}

// LoadConfig loads environment variables from the .env file
func LoadServerConfig() *ServerConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config := &ServerConfig{
		PORT:            getEnv("PORT", "8080"),
		ALLOWED_ORIGINS: getEnv("ALLOWED_ORIGINS", "*"),
	}

	return config
}
