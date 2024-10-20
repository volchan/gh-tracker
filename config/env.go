package config

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

var requiredEnv = []string{
	"APP_ENV",
	"PORT",
	"DB_HOST",
	"DB_PORT",
	"DB_USER",
	"DB_PASSWORD",
	"DB_NAME",
}

func loadEnvFile(path string, failIfMissing bool) {
	err := godotenv.Load(path)
	if err != nil && failIfMissing {
		log.Errorf("Error loading %s file", path)
		log.Fatal(err)
	}
}

func LoadEnv() {
	loadEnvFile(".env", true)
	loadEnvFile(fmt.Sprintf(".env.%s", Getenv("APP_ENV")), false)

	for _, key := range requiredEnv {
		if Getenv(key) == "" {
			log.Fatalf("Environment variable %s is required", key)
		}
	}
}

func Getenv(key string) string {
	return os.Getenv(key)
}
