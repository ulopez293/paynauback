package services

import (
	"context"
	"paynau-backend/models"
	"paynau-backend/prisma/db"
	"paynau-backend/repository"
)

func CreateProductoService(ctx context.Context, req models.CreateProductoRequest) (*db.ProductoModel, error) {
	return repository.CreateProducto(ctx, req.Nombre, req.Descripcion, req.Precio, req.Stock)
}

func GetAllProductosService(ctx context.Context) ([]db.ProductoModel, error) {
	return repository.GetProductos(ctx)
}

func GetProductoByIDService(ctx context.Context, id string) (*db.ProductoModel, error) {
	return repository.GetProductoByID(ctx, id)
}

func UpdateProductoService(ctx context.Context, id string, req models.UpdateProductoRequest) (*db.ProductoModel, error) {
	return repository.UpdateProducto(ctx, id, req.Nombre, req.Descripcion, req.Precio, req.Stock)
}

func DeleteProductoService(ctx context.Context, id string) error {
	return repository.DeleteProducto(ctx, id)
}
