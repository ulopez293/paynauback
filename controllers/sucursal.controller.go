package controllers

import (
	"context"
	"paynau-backend/models"
	"paynau-backend/services"

	"github.com/gofiber/fiber/v2"
)

func GetSucursal(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "GetUser ejecutado"})
}

func GetSucursales(c *fiber.Ctx) error {
	sucursales, err := services.GetSucursalesService(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error:   "Error obteniendo las sucursales",
			Details: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(sucursales)
}

func CreateSucursal(c *fiber.Ctx) error {
	var request models.CreateSucursalRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Formato de JSON inválido",
			Details: err.Error(),
		})
	}
	if request.Nombre == "" || request.Direccion == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "nombre y direccion son requeridos",
			Details: "",
		})
	}

	sucursal, err := services.CreateSucursalService(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error:   "Error creando la sucursal",
			Details: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(sucursal)
}

func UpdateSucursal(c *fiber.Ctx) error {
	id := c.Params("id")
	var request models.UpdateSucursalRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Formato JSON inválido",
			Details: err.Error(),
		})
	}

	if request.Nombre == "" || request.Direccion == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "nombre y direccion son requeridos",
			Details: "",
		})
	}

	sucursal, err := services.UpdateSucursalService(c.Context(), id, request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error:   "Error actualizando la sucursal",
			Details: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(sucursal)
}

func DeleteSucursal(c *fiber.Ctx) error {
	idStr := c.Params("id")

	err := services.DeleteSucursalService(context.Background(), idStr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al eliminar sucursal"})
	}

	return c.JSON(fiber.Map{"message": "Sucursal eliminada correctamente"})
}
