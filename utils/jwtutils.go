package utils

import (
	"errors"
	"fmt"
	"paynau-backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-that-no-one-knows")

func CreateJWT() (string, error) {
	claims := jwt.MapClaims{
		"sub":   "sub",
		"email": "email",
		"exp":   time.Now().Add(365 * 24 * time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func VerifyJWT(tokenStr string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid JWT")
}

func GetUserClaims(c *fiber.Ctx) (models.UserClaims, error) {
	raw := c.Locals("user")
	if raw == nil {
		return models.UserClaims{}, errors.New("no hay claims en el contexto")
	}
	claimsMap, ok := raw.(map[string]interface{})
	if !ok {
		return models.UserClaims{}, errors.New("formato de claims inválido")
	}
	subVal, ok := claimsMap["sub"].(string)
	if !ok {
		return models.UserClaims{}, errors.New("claim ‘sub’ no presente o no es string")
	}
	emailVal, ok := claimsMap["email"].(string)
	if !ok {
		return models.UserClaims{}, errors.New("claim ‘email’ no presente o no es string")
	}
	return models.UserClaims{
		Sub:   subVal,
		Email: emailVal,
	}, nil
}
