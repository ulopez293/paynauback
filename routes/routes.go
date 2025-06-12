package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API Paynau en ejecución")
	})

	api := app.Group("/api")
	ProductoRoutes(api)
	OrdenRoutes(api)
	AuthRoutes(api)
}
