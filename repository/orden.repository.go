package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"paynau-backend/models"
	"paynau-backend/prisma"
	"paynau-backend/prisma/db"
)

func CreateOrdenConProductos(ctx context.Context, req models.CreateOrdenRequest) (*db.OrdenModel, error) {
	client := prisma.GetPrisma()

	// 1) Validar stock y calcular total antes de la transacción
	var total float64
	productosInfo := make(map[string]*db.ProductoModel)

	for _, item := range req.Productos {
		p, err := client.Producto.FindUnique(db.Producto.ID.Equals(item.ProductoID)).Exec(ctx)
		if err != nil {
			return nil, err
		}
		if p == nil {
			return nil, errors.New("producto no encontrado: " + item.ProductoID)
		}
		if p.Stock < item.Cantidad {
			return nil, errors.New("stock insuficiente para producto: " + item.ProductoID)
		}
		productosInfo[item.ProductoID] = p
		total += float64(item.Cantidad) * p.Precio
	}

	// 2) Generar ID de orden
	ordenID := uuid.NewString()

	// 3) Construir todas las transacciones
	var txns []db.PrismaTransaction

	// Crear la orden (cabecera) con ID personalizado y total
	createOrden := client.Orden.CreateOne(
		db.Orden.Total.Set(total),
		db.Orden.ID.Set(ordenID),
	).Tx()
	txns = append(txns, createOrden)

	// Crear detalles y actualizar stock
	for _, item := range req.Productos {
		p := productosInfo[item.ProductoID]

		// Reducir stock
		txns = append(txns, client.Producto.FindUnique(
			db.Producto.ID.Equals(item.ProductoID),
		).Update(
			db.Producto.Stock.Set(p.Stock-item.Cantidad),
		).Tx())

		// Crear detalle vinculado a la orden
		txns = append(txns, client.OrdenProducto.CreateOne(
			db.OrdenProducto.Cantidad.Set(item.Cantidad),
			db.OrdenProducto.Total.Set(float64(item.Cantidad)*p.Precio),
			db.OrdenProducto.Orden.Link(db.Orden.ID.Equals(ordenID)),
			db.OrdenProducto.Producto.Link(db.Producto.ID.Equals(item.ProductoID)),
		).Tx())
	}

	// 4) Ejecutar la transacción atómica
	if err := client.Prisma.Transaction(txns...).Exec(ctx); err != nil {
		return nil, err
	}

	// 5) Recuperar la orden completa con sus detalles y productos anidados
	orden, err := client.Orden.FindUnique(
		db.Orden.ID.Equals(ordenID),
	).With(
		db.Orden.Productos.Fetch().With(db.OrdenProducto.Producto.Fetch()),
	).Exec(ctx)
	if err != nil || orden == nil {
		return nil, errors.New("no se pudo recuperar la orden creada")
	}

	return orden, nil
}
