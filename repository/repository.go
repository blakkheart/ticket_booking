package repository

import (
	"ticket-booking/domain/model"
)

type BaseRepository struct {
	//DB *sql.DB
}

func (repo *BaseRepository) Create(user model.User) (Account, error) {
	q := New(DB.Conn)
	account, err := q.CreateAccount(DB.Ctx,
		CreateAccountParams{
			Name:  user.Name,
			Email: user.Email,
		})

	return account, err
}

func (repo *BaseRepository) Delete(id int) (int, error) {
	var count int = 2
	return count, nil
}
func (repo *BaseRepository) Get(id int) (int, error) {
	var count int = 3
	return count, nil
}
func (repo *BaseRepository) GetMany(filter any) {

}
