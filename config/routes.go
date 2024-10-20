package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func (conf *Config) SetupRoutes(app *fiber.App) {
	app.Get("/metrics", monitor.New())
}
