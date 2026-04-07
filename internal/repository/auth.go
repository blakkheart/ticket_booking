package repository

import (
	"context"
	"ticket-booking/internal/auth/models"
)

type AuthRepository interface {
	Create(ctx context.Context, token *models.RefreshToken) (*models.RefreshToken, error)
	GetTokenByUserID(ctx context.Context, userID int64) (*models.RefreshToken, error)
	UpdateTokenByUserID(ctx context.Context, newtoken *models.RefreshToken) (*models.RefreshToken, error)
	GetTokenByHash(ctx context.Context, tokenHash string) (*models.RefreshToken, error)
}
