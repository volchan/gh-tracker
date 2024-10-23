package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/shareed2k/goth_fiber"
)

func GothRouter(app *fiber.App) {
	app.Get("/login/:provider", goth_fiber.BeginAuthHandler)

	app.Get("/auth/:provider/callback", func(ctx *fiber.Ctx) error {
		user, err := goth_fiber.CompleteUserAuth(ctx)
		if err != nil {
			log.Fatal(err)
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		return ctx.JSON(user)
	})

	app.Delete("/logout", func(ctx *fiber.Ctx) error {
		if err := goth_fiber.Logout(ctx); err != nil {
			log.Fatal(err)
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		return ctx.SendStatus(fiber.StatusOK)
	})
}
