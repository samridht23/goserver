package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samridht23/server/middleware"
)

func Initalize(app *fiber.App) {
	app.Use(middleware.Security)
	app.Use(middleware.Json)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hell World")
	})
}
