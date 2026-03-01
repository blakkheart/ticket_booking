package postgres

import (
	"context"
	"ticket-booking/internal/user"
	sqlc_repository "ticket-booking/repository/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type accountRepository struct {
	queries *sqlc_repository.Queries
}

func NewUserRepository(db *pgxpool.Pool) *accountRepository {
	return &accountRepository{
		queries: sqlc_repository.New(db),
	}
}

func (repo *accountRepository) Create(ctx context.Context, u *user.AccountIn) (*user.Account, error) {
	account, err := repo.queries.CreateAccount(
		ctx,
		sqlc_repository.CreateAccountParams{
			Name:     u.Name,
			Email:    u.Email,
			Password: u.Password,
			Role:     user.Member,
		},
	)

	return repo.fromSqlcAccount(&account), err
}

func (repo *accountRepository) Delete(id int64) error {
	return nil
}

func (repo *accountRepository) Get(id int64) (*user.Account, error) {
	return nil, nil
}

func (repo *accountRepository) GetMany(filter any) ([]*user.Account, error) {
	return nil, nil
}

func (repo *accountRepository) fromSqlcAccount(a *sqlc_repository.Account) *user.Account {
	account := &user.Account{
		ID:    a.ID,
		Email: a.Email,
		Name:  a.Name,
		Role:  a.Role,
	}
	return account
}
