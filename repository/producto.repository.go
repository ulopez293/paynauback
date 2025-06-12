package repository

import (
	"context"
	"paynau-backend/prisma"
	"paynau-backend/prisma/db"
)

func CreateProducto(ctx context.Context, nombre, descripcion string, precio float64, stock int) (*db.ProductoModel, error) {
	client := prisma.GetPrisma()
	return client.Producto.CreateOne(
		db.Producto.Nombre.Set(nombre),
		db.Producto.Descripcion.Set(descripcion),
		db.Producto.Precio.Set(precio),
		db.Producto.Stock.Set(stock),
	).Exec(ctx)
}

func GetProductos(ctx context.Context) ([]db.ProductoModel, error) {
	client := prisma.GetPrisma()
	return client.Producto.FindMany(
		db.Producto.Deleted.Equals(false), // solo los no eliminados
	).Exec(ctx)
}

func GetProductoByID(ctx context.Context, id string) (*db.ProductoModel, error) {
	client := prisma.GetPrisma()
	return client.Producto.FindUnique(
		db.Producto.ID.Equals(id),
	).Exec(ctx)
}

func UpdateProducto(ctx context.Context, id, nombre, descripcion string, precio float64, stock int) (*db.ProductoModel, error) {
	client := prisma.GetPrisma()
	return client.Producto.FindUnique(
		db.Producto.ID.Equals(id),
	).Update(
		db.Producto.Nombre.Set(nombre),
		db.Producto.Descripcion.Set(descripcion),
		db.Producto.Precio.Set(precio),
		db.Producto.Stock.Set(stock),
	).Exec(ctx)
}

func DeleteProducto(ctx context.Context, id string) error {
	client := prisma.GetPrisma()
	_, err := client.Producto.FindUnique(
		db.Producto.ID.Equals(id),
	).Update(
		db.Producto.Deleted.Set(true),
	).Exec(ctx)
	return err
}
