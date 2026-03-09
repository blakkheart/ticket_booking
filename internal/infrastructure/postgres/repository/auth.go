package repository

import (
	"context"
	"log/slog"
	"ticket-booking/internal/auth"
	sqlc_repository "ticket-booking/internal/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type authRepository struct {
	queries *sqlc_repository.Queries
	logger  *slog.Logger
}

func NewAuthRepository(db *pgxpool.Pool, logger *slog.Logger) *authRepository {
	return &authRepository{
		queries: sqlc_repository.New(db),
		logger:  logger,
	}
}

func (r *authRepository) Create(ctx context.Context, token *auth.RefreshToken) (*auth.RefreshToken, error) {
	t, err := r.queries.CreateToken(
		ctx,
		sqlc_repository.CreateTokenParams{
			UserID:     token.UserID,
			ExpiresAt:  token.ExpiresAt,
			Revoked:    token.Revoked,
			TokenHash:  token.TokenHash,
			ReplacedBy: &token.ReplacedBy,
		},
	)

	return r.fromSqlcAccount(&t), err
}

func (r *authRepository) GetTokenByUserID(ctx context.Context, userID int64) (*auth.RefreshToken, error) {
	token, err := r.queries.GetTokenByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return r.fromSqlcAccount(&token), nil
}

func (r *authRepository) GetTokenByHash(ctx context.Context, tokenHash string) (*auth.RefreshToken, error) {
	token, err := r.queries.GetTokenByHash(ctx, tokenHash)
	if err != nil {
		return nil, err
	}
	return r.fromSqlcAccount(&token), nil
}

func (r *authRepository) UpdateTokenByUserID(ctx context.Context, newToken *auth.RefreshToken) error {
	return r.queries.UpdateTokenByUserID(ctx, sqlc_repository.UpdateTokenByUserIDParams{
		UserID:     newToken.UserID,
		TokenHash:  newToken.TokenHash,
		ExpiresAt:  newToken.ExpiresAt,
		Revoked:    newToken.Revoked,
		ReplacedBy: &newToken.ReplacedBy,
	})
}

func (r *authRepository) fromSqlcAccount(rt *sqlc_repository.RefreshToken) *auth.RefreshToken {
	account := &auth.RefreshToken{
		UserID:     rt.UserID,
		ExpiresAt:  rt.ExpiresAt,
		Revoked:    rt.Revoked,
		TokenHash:  rt.TokenHash,
		ReplacedBy: *rt.ReplacedBy,
	}
	return account
}
