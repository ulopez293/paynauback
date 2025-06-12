package repository

import (
	"context"
	"paynau-backend/prisma"
	"paynau-backend/prisma/db"
)

func FindUserByEmail(ctx context.Context, email string) (*db.UserModel, error) {
	client := prisma.GetPrisma()
	return client.User.FindUnique(
		db.User.Email.Equals(email),
	).Exec(ctx)
}

func SaveUser(ctx context.Context, user *db.UserModel) error {
	client := prisma.GetPrisma()
	_, err := client.User.
		FindUnique(db.User.IDUser.Equals(user.IDUser)).
		Update(
			db.User.Email.Set(user.Email),
			db.User.Password.Set(user.Password),
		).
		Exec(ctx)
	return err
}

func CreateUser(ctx context.Context, email, hashedPassword, rol, sucursalID string) (*db.UserModel, error) {
	client := prisma.GetPrisma()
	return client.User.CreateOne(
		db.User.Email.Set(email),
		db.User.Password.Set(hashedPassword),
		db.User.RolUser.Set(rol),
		db.User.SucursalID.Set(sucursalID), // obligatoria
	).Exec(ctx)
}

func ListUsers(ctx context.Context) ([]db.UserModel, error) {
	client := prisma.GetPrisma()
	return client.User.
		FindMany(
			db.User.RolUser.NotIn([]string{"ADMINISTRADOR"}),
		).
		Exec(ctx)
}

func DeleteUser(ctx context.Context, id string) error {
	client := prisma.GetPrisma()
	_, err := client.User.
		FindUnique(
			db.User.IDUser.Equals(id),
		).
		Delete().
		Exec(ctx)
	return err
}
