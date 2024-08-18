package config

import (
	"log"

	"github.com/joho/godotenv"
)

// Config struct to hold all configuration variables
type DbConfig struct {
	HOST     string
	PORT     string
	USER     string
	DATABASE string
	PASSWORD string
}

// LoadConfig loads environment variables from the .env file
func LoadDbConfig() *DbConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading db config from .env file: %v", err)
	}

	config := &DbConfig{
		HOST:     getEnv("DB_HOST", ""),
		PORT:     getEnv("DB_PORT", ""),
		USER:     getEnv("DB_USER", ""),
		DATABASE: getEnv("DB_NAME", ""),
		PASSWORD: getEnv("DB_PASSWORD", ""),
	}

	return config
}
