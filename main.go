package main

import (
	"fmt"
	"gh-tracker/config"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.LoadEnv()

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("could not load database")
	}

	c := &config.Config{
		DB: db,
	}

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(compress.New())
	app.Use(helmet.New())
	app.Use(idempotency.New())
	app.Use(logger.New())

	c.SetupRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Getenv("PORT"))))
}
