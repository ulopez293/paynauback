package controllers

import (
	"paynau-backend/models"
	"paynau-backend/services"

	"github.com/gofiber/fiber/v2"
)

func GetProductos(c *fiber.Ctx) error {
	productos, err := services.GetAllProductosService(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Error al obtener productos",
			"detail": err.Error(),
		})
	}
	return c.JSON(productos)
}

func GetProductoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID de producto es requerido"})
	}
	producto, err := services.GetProductoByIDService(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "No se pudo obtener producto", "detail": err.Error()})
	}
	if producto == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Producto no encontrado"})
	}
	return c.JSON(producto)
}

func CreateProducto(c *fiber.Ctx) error {
	var req models.CreateProductoRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Formato inválido"})
	}
	producto, err := services.CreateProductoService(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "No se pudo crear producto", "detail": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(producto)
}

func UpdateProducto(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID de producto es requerido"})
	}
	var req models.UpdateProductoRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Formato inválido"})
	}
	producto, err := services.UpdateProductoService(c.Context(), id, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "No se pudo actualizar producto", "detail": err.Error()})
	}
	return c.JSON(producto)
}

func DeleteProducto(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID de producto es requerido"})
	}
	if err := services.DeleteProductoService(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
