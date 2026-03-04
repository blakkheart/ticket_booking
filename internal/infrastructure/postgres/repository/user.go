package repository

import (
	"context"
	"errors"
	"log/slog"
	sqlc_repository "ticket-booking/internal/db/sqlc"
	"ticket-booking/internal/user"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type accountRepository struct {
	queries *sqlc_repository.Queries
	logger  *slog.Logger
}

func NewUserRepository(db *pgxpool.Pool, logger *slog.Logger) *accountRepository {
	return &accountRepository{
		queries: sqlc_repository.New(db),
		logger:  logger,
	}
}

func (repo *accountRepository) Create(ctx context.Context, u *user.CreateUserRequest) (*user.User, error) {
	account, err := repo.queries.CreateAccount(
		ctx,
		sqlc_repository.CreateAccountParams{
			Name:     u.Name,
			Email:    u.Email,
			Password: u.Password,
			Role:     user.Member,
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

func (repo *accountRepository) Delete(id int64) error {
	return nil
}

func (repo *accountRepository) Get(id int64) (*user.User, error) {
	return nil, nil
}

func (repo *accountRepository) GetMany(filter any) ([]*user.User, error) {
	return nil, nil
}

func (repo *accountRepository) fromSqlcAccount(a *sqlc_repository.Account) *user.User {
	account := &user.User{
		ID:    a.ID,
		Email: a.Email,
		Name:  a.Name,
		Role:  a.Role,
	}
	return account
}
