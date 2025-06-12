package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashPassword recibe una contraseña en texto plano y devuelve el hash SHA-256 en hexadecimal.
func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

// ComparePassword recibe la contraseña en texto plano y un hash previamente almacenado,
// genera el hash de la contraseña y lo compara con el hash dado.
// Devuelve true si coinciden.
func ComparePassword(password, hashed string) bool {
	return HashPassword(password) == hashed
}
