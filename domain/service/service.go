package service

import (
	"log"
	"time"

	"ticket-booking/domain/model"
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

func (service *Service) GetUser() model.User {
	return model.User{
		ID:    1,
		Name:  "Name",
		Email: "Email",
	}
}

func (service *Service) CreateAccount(user model.User) repository.Account {
	value, err := service.Repo.Create(user)
	if err != nil {
		log.Fatal("Something wrong with repo")
	}
	return value
}

func (service *Service) GetEvent() model.Event {
	return model.Event{
		ID:          1,
		Title:       "Concert",
		Description: "Musical",
		Date:        time.Date(2025, time.April, 15, 20, 0, 0, 0, time.UTC),
		Location:    "New York",
		TicketTypes: []model.TicketType{
			{
				ID:          1,
				Name:        "Common",
				Description: "Main event",
				Price:       50.0,
				Quantity:    200,
				EventID:     1,
			},
			{
				ID:          2,
				Name:        "VIP",
				Description: "Main Event + Autograph",
				Price:       100,
				Quantity:    100,
				EventID:     1,
			},
		},
	}
}

func (service *Service) GetBooking(user model.User, event model.Event) model.Booking {
	return model.Booking{
		ID:           1,
		UserID:       user.ID,
		EventID:      event.ID,
		TicketTypeID: event.TicketTypes[0].ID,
		Quantity:     2,
	}
}
