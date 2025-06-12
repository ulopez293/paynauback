// middlewares/protected.go
package middlewares

import (
	"strings"

	"paynau-backend/models"
	"paynau-backend/utils"

	"github.com/gofiber/fiber/v2"
)

func Protected(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: "Token requerido",
		})
	}
	token := parts[1]
	payload, err := utils.VerifyJWT(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error:   "Token inválido o expirado",
			Details: err.Error(),
		})
	}
	c.Locals("user", payload)
	return c.Next()
}
