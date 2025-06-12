package services

import (
	"context"
	"paynau-backend/models"
	"paynau-backend/utils"
)

func AuthUserService(ctx context.Context) (*models.AuthResponse, error) {
	token, err := utils.CreateJWT()
	if err != nil {
		return nil, err
	}
	return &models.AuthResponse{
		Token: token,
	}, nil
}
