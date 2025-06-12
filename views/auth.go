package views

import (
	"paynau-backend/models"

	"github.com/gofiber/fiber/v2"
)

func JSONLoginSuccess(c *fiber.Ctx, token string) error {
	return c.Status(fiber.StatusOK).JSON(models.AuthResponse{
		Token: token,
	})
}

func JSONAuthError(c *fiber.Ctx, status int, errMsg string, details string) error {
	return c.Status(status).JSON(models.ErrorResponse{
		Error:   errMsg,
		Details: details,
	})
}
