package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func (conf *Config) SetupRoutes(app *fiber.App) {
	app.Get("/metrics", monitor.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		conf.DB.Exec("SELECT 1")
		return ctx.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})
}
