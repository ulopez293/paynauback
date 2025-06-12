package routes

import (
	"paynau-backend/controllers"
	"paynau-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ProductoRoutes(app fiber.Router) {
	products := app.Group("/products")

	// GET /api/products
	products.Get("/", controllers.GetProductos)
	// POST /api/products
	products.Post("/", middlewares.Protected, controllers.CreateProducto)
	// PUT /api/products/:id
	products.Put("/:id", middlewares.Protected, controllers.UpdateProducto)
	// DELETE /api/products/:id
	products.Delete("/:id", middlewares.Protected, controllers.DeleteProducto)
}
