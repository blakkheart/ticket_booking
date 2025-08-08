package model_interface

import (
	"ticket-booking/repository"
)

type BaseRepository interface {
	Create(user repository.Account) (repository.Account, error)
	Delete(id int) (int, error)
	Get(id int) (int, error)
	GetMany(any)
	CreateEvent(repository.CreateEventParams) (repository.Event, error)
	GetEvent(id int64) (repository.Event, error)
}
