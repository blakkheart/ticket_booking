package model_interface

import (
	"ticket-booking/repository"
)

type IAccountRepository interface {
	Create(user repository.CreateAccountParams) (repository.Account, error)
	Delete(id int) (int, error)
	Get(id int) (int, error)
	GetMany(any)
}
