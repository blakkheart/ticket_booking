package repository

import (
	"context"
	"errors"
	"log/slog"
	sqlc_repository "ticket-booking/internal/db/sqlc"
	"ticket-booking/internal/user"
	"ticket-booking/internal/user/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type accountRepository struct {
	queries *sqlc_repository.Queries
	logger  *slog.Logger
}

func NewUserRepository(
	db *pgxpool.Pool,
	logger *slog.Logger,
) *accountRepository {
	return &accountRepository{
		queries: sqlc_repository.New(db),
		logger:  logger,
	}
}

func (repo *accountRepository) Create(
	ctx context.Context,
	u *models.CreateUserParams,
) (*models.User, error) {
	account, err := repo.queries.CreateAccount(
		ctx,
		sqlc_repository.CreateAccountParams{
			ID:       u.ID,
			Name:     u.Name,
			Email:    u.Email,
			Password: u.Password,
			Role:     u.Role,
		},
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" && pgErr.ConstraintName == "account_email_key" {
				return nil, user.ErrEmailAlreadyUsed
			}
		}
	}

	return repo.fromSqlcAccount(&account), err
}

func (repo *accountRepository) Delete(id uuid.UUID) error {
	return nil
}

func (repo *accountRepository) Get(id uuid.UUID) (*models.User, error) {
	return nil, nil
}

func (repo *accountRepository) GetMany(filter any) ([]*models.User, error) {
	return nil, nil
}

func (repo *accountRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*models.UserAuth, error) {
	acc, err := repo.queries.GetAccountByEmail(ctx, email)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return nil, err
		}
	}
	return &models.UserAuth{
		ID:           acc.ID,
		Email:        acc.Email,
		PasswordHash: acc.Password,
		Role:         acc.Role,
	}, nil

}

func (repo *accountRepository) fromSqlcAccount(
	a *sqlc_repository.Account,
) *models.User {
	account := &models.User{
		ID:    a.ID,
		Email: a.Email,
		Name:  a.Name,
		Role:  a.Role,
	}
	return account
}
