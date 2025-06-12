package services

import (
	"context"
	"fmt"
	"paynau-backend/models"
	"paynau-backend/prisma/db"
	"paynau-backend/repository"
	"paynau-backend/utils"
	"strings"
)

func GetUserByEmail(ctx context.Context, email string) (*db.UserModel, error) {
	return repository.FindUserByEmail(ctx, email)
}

func UpdateUserService(ctx context.Context, oldEmail string, req models.UpdateUserRequest) (*db.UserModel, error) {
	// 1. Buscar por email antiguo
	user, err := repository.FindUserByEmail(ctx, oldEmail)
	if err != nil || user == nil {
		return nil, err
	}

	// 2. Si incluyen nueva contraseña, hashearla
	if req.Password != "" {
		user.Password = utils.HashPassword(req.Password)
	}

	// 3. Actualizar al nuevo email
	lowerEmail := strings.ToLower(req.Email)
	user.Email = lowerEmail

	// 4. Guardar cambios
	if err := repository.SaveUser(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUserService(ctx context.Context, req models.CreateUserRequest) (*db.UserModel, error) {
	hashed := utils.HashPassword(req.Password)
	return repository.CreateUser(ctx, req.Email, hashed, req.RolUser, req.SucursalID)
}

func GetUsersService(ctx context.Context) ([]db.UserModel, error) {
	return repository.ListUsers(ctx)
}

func DeleteUserService(ctx context.Context, id string) error {
	// aquí podrías comprobar permisos, o no dejar borrar al propio administrador, etc.
	err := repository.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("no se pudo eliminar el usuario: %w", err)
	}
	return nil
}
