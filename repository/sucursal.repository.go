package repository

import (
	"context"
	"paynau-backend/prisma"
	"paynau-backend/prisma/db"
)

func CreateSucursal(ctx context.Context, nombre, direccion string) (*db.SucursalModel, error) {
	client := prisma.GetPrisma()

	sucursal, err := client.Sucursal.CreateOne(
		db.Sucursal.Nombre.Set(nombre),
		db.Sucursal.Direccion.Set(direccion),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return sucursal, nil
}

func GetSucursales(ctx context.Context) ([]db.SucursalModel, error) {
	client := prisma.GetPrisma()

	sucursales, err := client.Sucursal.FindMany().Exec(ctx)
	if err != nil {
		return nil, err
	}

	return sucursales, nil
}

func DeleteSucursal(ctx context.Context, id string) error {
	client := prisma.GetPrisma()

	_, err := client.Sucursal.FindUnique(
		db.Sucursal.ID.Equals(id),
	).Delete().Exec(ctx)

	return err
}

func HasUsersInSucursal(ctx context.Context, sucursalId string) (bool, error) {
	client := prisma.GetPrisma()

	users, err := client.User.FindMany(
		db.User.SucursalID.Equals(sucursalId),
	).Exec(ctx)
	if err != nil {
		return false, err
	}

	return len(users) > 0, nil
}

func UpdateSucursal(ctx context.Context, id string, nombre string, direccion string) (*db.SucursalModel, error) {
	client := prisma.GetPrisma()

	return client.Sucursal.FindUnique(
		db.Sucursal.ID.Equals(id),
	).Update(
		db.Sucursal.Nombre.Set(nombre),
		db.Sucursal.Direccion.Set(direccion),
	).Exec(ctx)
}
