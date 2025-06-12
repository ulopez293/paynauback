package routes

import (
	"paynau-backend/controllers"
	"paynau-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SucursalRoutes(app fiber.Router) {
	sucursales := app.Group("/sucursales")

	sucursales.Post("/sucursal", middlewares.Protected, controllers.CreateSucursal)
	sucursales.Get("/sucursales", controllers.GetSucursales)
	sucursales.Delete("/sucursal/:id", middlewares.Protected, controllers.DeleteSucursal)
	sucursales.Put("/sucursal/:id", middlewares.Protected, controllers.UpdateSucursal)
}
