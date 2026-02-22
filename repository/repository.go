package repository

import (
	"ticket-booking/models/database"
	accountModel "ticket-booking/models/domain/account"
)

type BaseRepository struct {
	DB *database.DBStruct
}

func (repo *BaseRepository) Create(user accountModel.AccountIn) (Account, error) {
	q := New(repo.DB.Conn)
	account, err := q.CreateAccount(repo.DB.Ctx,
		CreateAccountParams{
			Name:  user.Name,
			Email: user.Email,
		})

	return account, err
}

func (repo *BaseRepository) Delete(id int64) error {
	return nil
}

func (repo *BaseRepository) Get(id int64) (Account, error) {
	return Account{}, nil
}

func (repo *BaseRepository) GetMany(filter any) ([]Account, error) {
	return []Account{}, nil

}

func (repo *BaseRepository) CreateEvent(event_params CreateEventParams) (Event, error) {
	q := New(repo.DB.Conn)
	event, err := q.CreateEvent(
		repo.DB.Ctx,
		event_params,
	)
	return event, err
}

func (repo *BaseRepository) GetEvent(id int64) (Event, error) {
	q := New(repo.DB.Conn)
	event, err := q.GetEvent(
		repo.DB.Ctx,
		id,
	)
	return event, err
}
