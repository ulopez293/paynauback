// routes/users.go
package routes

import (
	"paynau-backend/controllers"
	"paynau-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app fiber.Router) {
	users := app.Group("/users")

	users.Post("/", middlewares.Protected, controllers.CreateUser)
	users.Post("/login", controllers.LoginUser)
	users.Get("/", middlewares.Protected, controllers.GetUser)
	users.Get("/all", middlewares.Protected, controllers.GetUsers)
	users.Put("/", middlewares.Protected, controllers.UpdateUser)
	users.Delete("/:id", middlewares.Protected, controllers.DeleteUser)
}
