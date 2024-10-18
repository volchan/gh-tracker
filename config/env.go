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

func LoadEnv() {
	godotenv.Load(".env")
	godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("APP_ENV")))

	for _, key := range requiredEnv {
		if Getenv(key) == "" {
			log.Fatalf("Environment variable %s is required", key)
		}
	}
}

func Getenv(key string) string {
	return os.Getenv(key)
}
