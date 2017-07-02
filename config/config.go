package config

import (
	"log"
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
)

// TODO: use the viper package for configuration instead

var configLoaded = false

// Get reads returns a configuration value for the given key.
// If the .env file has not been read, it will load it.
func Get(key string) string {
	if !configLoaded {
		var _, filename, _, ok = runtime.Caller(0)
		if !ok {
			panic("No caller information")
		}
		var packageDir = path.Dir(filename)

		var err = godotenv.Load(packageDir + "/../.env")
		if err != nil {
			log.Fatal(err)
		}
		configLoaded = true
	}

	switch key {
	case "DevPort":
		return os.Getenv("DEV_PORT")
	case "Environment":
		return os.Getenv("ENVIRONMENT")
	default:
		panic("Environment variable " + key + " not found!")
	}
}
