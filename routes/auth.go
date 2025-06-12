package routes

import (
	"paynau-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router) {
	router.Post("/auth", controllers.Auth)
}
