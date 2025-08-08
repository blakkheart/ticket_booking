package service

import (
	"log"

	model_interface "ticket-booking/domain/model/interface"
	"ticket-booking/repository"
)

type Service struct {
	Repo model_interface.BaseRepository
}

func (service *Service) Get(id int) int {
	value, err := service.Repo.Get(id)
	if err != nil {
		log.Fatal("Something wrong with repo")
	}
	return value
}

func (service *Service) GetMany(filters any) {
}

func (service *Service) GetUser() repository.Account {
	return repository.Account{
		ID:    1,
		Name:  "Name",
		Email: "Email",
	}
}

func (service *Service) CreateAccount(user repository.Account) repository.Account {
	value, err := service.Repo.Create(user)
	if err != nil {
		log.Fatal("Something wrong with repo")
	}
	return value
}

func (service *Service) GetEvent(id int64) repository.Event {
	//event_date := time.Date(2025, time.April, 15, 20, 0, 0, 0, time.UTC)
	event, err := service.Repo.GetEvent(id)
	if err != nil {
		log.Fatal("Something wrong with repo")
	}
	return event
}

func (service *Service) CreateEvent(event_param repository.CreateEventParams) repository.Event {
	value, err := service.Repo.CreateEvent(event_param)
	if err != nil {
		log.Fatal("Something wrong with repo")
	}
	return value
}

func (service *Service) GetBooking(user repository.Account, event repository.Event) repository.Booking {
	return repository.Booking{
		ID:        1,
		AccountID: user.ID,
		EventID:   event.ID,
		Quantity:  2,
	}
}
