package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// LoadConfig: load the config based on .env file or on local env automatically parse to struct.
func LoadConfig(config interface{}) (err error) {
	_ = godotenv.Load(".env")

	if err := env.Parse(config); err != nil {
		log.Printf("Failed to parse environment variables: %v", err)
		return err
	}

	return nil
}
