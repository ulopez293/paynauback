package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API Pollería en ejecución")
	})

	api := app.Group("/api")
	ProductoRoutes(api)
}
