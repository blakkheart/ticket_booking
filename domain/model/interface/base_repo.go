package model_interface

import (
	"ticket-booking/domain/model"
	"ticket-booking/repository"
)

type BaseRepository interface {
	Create(user model.User) (repository.Account, error)
	Delete(id int) (int, error)
	Get(id int) (int, error)
	GetMany(any)
}
