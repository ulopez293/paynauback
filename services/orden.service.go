package services

import (
	"context"
	"paynau-backend/models"
	"paynau-backend/prisma/db"
	"paynau-backend/repository"
)

func CreateOrdenService(ctx context.Context, req models.CreateOrdenRequest) (*db.OrdenModel, error) {
	return repository.CreateOrdenConProductos(ctx, req)
}

func GetOrdenesService(ctx context.Context) ([]db.OrdenModel, error) {
	return repository.ObtenerOrdenesConProductos(ctx)
}
