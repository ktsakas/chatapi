package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	// Port for the app in dev mode
	DevPort string
}

var cachedConfig *config

// GetConfig reads the configuration from the .env file
// into the variables exported by the config package
// if it has already been read it returns a cached version
func GetConfig() *config {
	if cachedConfig == nil {
		var err = godotenv.Load(".env")
		if err != nil {
			log.Fatal(err)
		}

		cachedConfig = &config{
			DevPort: os.Getenv("DEV_PORT"),
		}
	}

	return cachedConfig
}
