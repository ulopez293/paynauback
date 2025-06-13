package utils_test

import (
	"paynau-backend/utils"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "miSuperSecreto123"
	hashed := utils.HashPassword(password)

	if hashed == "" {
		t.Error("HashPassword retornó cadena vacía")
	}

	if hashed == password {
		t.Error("HashPassword no debería devolver la misma cadena que la contraseña")
	}
}

func TestComparePassword(t *testing.T) {
	password := "otraClave456"
	hashed := utils.HashPassword(password)

	if !utils.ComparePassword(password, hashed) {
		t.Error("ComparePassword debería retornar true para la contraseña y hash correspondientes")
	}

	wrongPassword := "claveIncorrecta"
	if utils.ComparePassword(wrongPassword, hashed) {
		t.Error("ComparePassword debería retornar false para una contraseña incorrecta")
	}
}
