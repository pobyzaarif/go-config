package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// LoadConfig: load the config based on .env file or on local env automatically parse to struct.
func LoadConfig(config interface{}) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("No .env file found or failed to load: %v", err)
	}

	if err := env.Parse(config); err != nil {
		log.Printf("Failed to parse environment variables: %v", err)
	}
}
