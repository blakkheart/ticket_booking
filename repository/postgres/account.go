package postgres

import (
	"ticket-booking/domain/model"
	"ticket-booking/models/database"
	"ticket-booking/repository"
)

type BaseRepository struct {
	DB *database.DBStruct
}

func (repo *BaseRepository) Create(user *model.AccountIn) (*model.Account, error) {
	q := repository.New(repo.DB.Conn)
	account, err := q.CreateAccount(repo.DB.Ctx,
		repository.CreateAccountParams{
			Name:  user.Name,
			Email: user.Email,
		})

	return repo.fromSqlcAccount(&account), err
}

func (repo *BaseRepository) fromSqlcAccount(a *repository.Account) *model.Account {
	account := &model.Account{
		ID:    a.ID,
		Email: a.Email,
		Name:  a.Name,
		Role:  a.Role,
	}
	return account
}
