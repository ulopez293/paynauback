package controllers

import (
	"context"
	"paynau-backend/models"
	"paynau-backend/services"

	"github.com/gofiber/fiber/v2"
)

func CreateOrden(c *fiber.Ctx) error {
	var req models.CreateOrdenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	order, err := services.CreateOrdenService(context.Background(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}

func GetOrdenes(c *fiber.Ctx) error {
	ordenes, err := services.GetOrdenesService(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(ordenes)
}
