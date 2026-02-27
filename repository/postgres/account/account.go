package accrep

import (
	"context"
	accountModel "ticket-booking/models/domain/account"
	sqlc_repository "ticket-booking/repository/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type accountRepository struct {
	DB      *pgxpool.Pool
	queries *sqlc_repository.Queries
}

func New(db *pgxpool.Pool) *accountRepository {
	return &accountRepository{
		DB:      db,
		queries: sqlc_repository.New(db),
	}
}

func (repo *accountRepository) Create(ctx context.Context, user *accountModel.AccountIn) (*accountModel.Account, error) {
	account, err := repo.queries.CreateAccount(
		ctx,
		sqlc_repository.CreateAccountParams{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
			Role:     accountModel.Member,
		},
	)

	return repo.fromSqlcAccount(&account), err
}

func (repo *accountRepository) Delete(id int64) error {
	return nil
}

func (repo *accountRepository) Get(id int64) (*accountModel.Account, error) {
	return nil, nil
}

func (repo *accountRepository) GetMany(filter any) ([]*accountModel.Account, error) {
	return nil, nil
}

func (repo *accountRepository) fromSqlcAccount(a *sqlc_repository.Account) *accountModel.Account {
	account := &accountModel.Account{
		ID:    a.ID,
		Email: a.Email,
		Name:  a.Name,
		Role:  a.Role,
	}
	return account
}
