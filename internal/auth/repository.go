package auth

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, token *RefreshToken) (*RefreshToken, error)
	GetTokenByUserID(ctx context.Context, userID int64) (*RefreshToken, error)
	UpdateTokenByUserID(ctx context.Context, newtoken *RefreshToken) error
}
