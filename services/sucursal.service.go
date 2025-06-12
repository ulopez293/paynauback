package services

import (
	"context"
	"fmt"
	"paynau-backend/models"
	"paynau-backend/prisma/db"
	"paynau-backend/repository"
)

func GetSucursalesService(ctx context.Context) ([]db.SucursalModel, error) {
	return repository.GetSucursales(ctx)
}

func CreateSucursalService(ctx context.Context, req models.CreateSucursalRequest) (*db.SucursalModel, error) {
	return repository.CreateSucursal(ctx, req.Nombre, req.Direccion)
}

func DeleteSucursalService(ctx context.Context, id string) error {
	hasUsers, err := repository.HasUsersInSucursal(ctx, id)
	if err != nil {
		return err
	}

	if hasUsers {
		return fmt.Errorf("no se puede eliminar la sucursal porque tiene usuarios asociados")
	}

	return repository.DeleteSucursal(ctx, id)
}

func UpdateSucursalService(ctx context.Context, id string, req models.UpdateSucursalRequest) (*db.SucursalModel, error) {
	return repository.UpdateSucursal(ctx, id, req.Nombre, req.Direccion)
}
