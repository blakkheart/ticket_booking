package postgres

import (
	"ticket-booking/models/database"
	accountModel "ticket-booking/models/domain/account"
	"ticket-booking/repository"
)

type BaseRepository struct {
	DB *database.DBStruct
}

func (repo *BaseRepository) Create(user *accountModel.AccountIn) (*accountModel.Account, error) {
	q := repository.New(repo.DB.Conn)
	account, err := q.CreateAccount(repo.DB.Ctx,
		repository.CreateAccountParams{
			Name:  user.Name,
			Email: user.Email,
		})

	return repo.fromSqlcAccount(&account), err
}

func (repo *BaseRepository) fromSqlcAccount(a *repository.Account) *accountModel.Account {
	account := &accountModel.Account{
		ID:    a.ID,
		Email: a.Email,
		Name:  a.Name,
		Role:  a.Role,
	}
	return account
}
