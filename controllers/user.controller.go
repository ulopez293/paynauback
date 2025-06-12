package controllers

import (
	"paynau-backend/models"
	"paynau-backend/services"
	"paynau-backend/utils"
	"paynau-backend/views"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	email := strings.ToLower(c.Query("email"))
	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email es requerido",
		})
	}
	user, err := services.GetUserByEmail(c.Context(), email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener el usuario",
		})
	}
	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Usuario no encontrado",
		})
	}
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Formato inválido"})
	}
	if req.Email == "" || req.Password == "" || req.RolUser == "" || req.SucursalID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email, contraseña, rol y sucursal son requeridos"})
	}
	user, err := services.CreateUserService(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "No se pudo crear usuario", "detail": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func LoginUser(c *fiber.Ctx) error {
	var request models.LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return views.JSONAuthError(c, fiber.StatusBadRequest, "Formato de JSON inválido", err.Error())
	}
	request.Email = strings.ToLower(request.Email)
	if request.Email == "" || request.Password == "" {
		return views.JSONAuthError(c, fiber.StatusBadRequest, "Email y contraseña son requeridos", "")
	}
	userTokenPair, err := services.AuthUserService(c.UserContext(), request.Email, request.Password)
	if err != nil {
		return views.JSONAuthError(c, fiber.StatusBadRequest, "Error en autenticación", err.Error())
	}
	return views.JSONLoginSuccess(c, userTokenPair.User, userTokenPair.Token)
}

func UpdateUser(c *fiber.Ctx) error {
	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Formato inválido"})
	}
	if req.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email es requerido"})
	}
	claims, err := utils.GetUserClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "No autorizado"})
	}
	// Actualiza al usuario y recibe el modelo prisma
	updatedUser, err := services.UpdateUserService(c.Context(), claims.Email, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "No se pudo actualizar", "detail": err.Error()})
	}
	// Genera nuevo token usando el helper
	newToken, err := services.SignUserToken(updatedUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "No se pudo crear token"})
	}
	// Devuelve usuario y token
	return c.JSON(fiber.Map{
		"user":  updatedUser,
		"token": newToken,
	})
}

func GetUsers(c *fiber.Ctx) error {
	users, err := services.GetUsersService(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error obteniendo usuarios",
			"details": err.Error(),
		})
	}
	return c.JSON(users)
}

func DeleteUser(c *fiber.Ctx) error {
	// obtenemos el id de la ruta: DELETE /api/users/:id
	id := c.Params("id")
	if id == "" {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "ID de usuario es requerido"})
	}
	if err := services.DeleteUserService(c.Context(), id); err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}
	// 204 No Content
	return c.SendStatus(fiber.StatusNoContent)
}
