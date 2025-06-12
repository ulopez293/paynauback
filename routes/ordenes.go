package routes

import (
	"paynau-backend/controllers"
	"paynau-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func OrdenRoutes(app fiber.Router) {
	orders := app.Group("/orders")
	orders.Post("/", middlewares.Protected, controllers.CreateOrden)
	orders.Get("/", middlewares.Protected, controllers.GetOrdenes)
}
