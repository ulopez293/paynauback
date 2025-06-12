package controllers

import (
	"paynau-backend/services"
	"paynau-backend/views"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	userTokenPair, err := services.AuthUserService(c.UserContext())
	if err != nil {
		return views.JSONAuthError(c, fiber.StatusBadRequest, "Error en autenticación", err.Error())
	}
	return views.JSONAuthSuccess(c, userTokenPair.Token)
}
