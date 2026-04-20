package repository

import (
	"context"
	"errors"
	"log/slog"
	"ticket-booking/internal/auth"
	"ticket-booking/internal/auth/models"
	sqlc_repository "ticket-booking/internal/db/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type authRepository struct {
	queries *sqlc_repository.Queries
	logger  *slog.Logger
}

func NewAuthRepository(
	db *pgxpool.Pool,
	logger *slog.Logger,
) *authRepository {
	return &authRepository{
		queries: sqlc_repository.New(db),
		logger:  logger,
	}
}

func (r *authRepository) Create(
	ctx context.Context,
	token *models.RefreshToken,
) (*models.RefreshToken, error) {
	t, err := r.queries.CreateToken(
		ctx,
		sqlc_repository.CreateTokenParams{
			UserID:     token.UserID,
			ExpiresAt:  token.ExpiresAt,
			Revoked:    token.Revoked,
			TokenHash:  token.TokenHash,
			ReplacedBy: token.ReplacedBy,
		},
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" && pgErr.ConstraintName == "account_email_key" {
				return nil, auth.RefreshTokenDuplicationError
			}
		}
	}

	return r.fromSqlcAuth(&t), err
}

func (r *authRepository) GetTokenByUserID(
	ctx context.Context,
	userID uuid.UUID,
) (*models.RefreshToken, error) {
	token, err := r.queries.GetTokenByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return r.fromSqlcAuth(&token), nil
}

func (r *authRepository) GetTokenByHash(
	ctx context.Context,
	tokenHash string,
) (*models.RefreshToken, error) {
	token, err := r.queries.GetTokenByHash(ctx, tokenHash)
	if err != nil {
		return nil, err
	}
	return r.fromSqlcAuth(&token), nil
}

func (r *authRepository) UpdateTokenByUserID(
	ctx context.Context,
	newToken *models.RefreshToken,
) (*models.RefreshToken, error) {
	token, err := r.queries.UpdateTokenByUserID(ctx, sqlc_repository.UpdateTokenByUserIDParams{
		UserID:     newToken.UserID,
		TokenHash:  newToken.TokenHash,
		ExpiresAt:  newToken.ExpiresAt,
		Revoked:    newToken.Revoked,
		ReplacedBy: newToken.ReplacedBy,
	})
	if err != nil {
		return nil, err
	}
	return r.fromSqlcAuth(&token), nil
}

func (r *authRepository) fromSqlcAuth(
	rt *sqlc_repository.RefreshToken,
) *models.RefreshToken {
	refreshToken := &models.RefreshToken{
		UserID:     rt.UserID,
		ExpiresAt:  rt.ExpiresAt,
		Revoked:    rt.Revoked,
		TokenHash:  rt.TokenHash,
		ReplacedBy: rt.ReplacedBy,
	}
	return refreshToken
}
