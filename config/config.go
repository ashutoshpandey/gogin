package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct to hold all configuration variables
type Config struct {
	PORT string
}

// LoadConfig loads environment variables from the .env file
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config := &Config{
		PORT: getEnv("PORT", "8080"),
	}

	return config
}

// getEnv is a helper function to read an environment variable or return a default value
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
