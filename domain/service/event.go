package service

import (
	"context"
	"log"

	model "ticket-booking/models/domain/event"
	model_interface "ticket-booking/models/domain/interface"
)

type EventService interface {
	Get(id int64) *model.Event
	GetMany(filters any) []*model.Event
	Create(ctx context.Context, user *model.Event) *model.Event
}

func NewE(repo model_interface.IEventRepository) EventService {
	return &eventService{Repo: repo}
}

type eventService struct {
	Repo model_interface.IEventRepository
}

func (service *eventService) Get(id int64) *model.Event {
	value, err := service.Repo.Get(id)
	if err != nil {
		log.Fatal("Something wrong with repo")
	}
	return value
}

func (service *eventService) GetMany(filters any) []*model.Event {
	return nil
}

func (service *eventService) Create(ctx context.Context, user *model.Event) *model.Event {

	u := &model.Event{
		ID: 1,
	}

	value, err := service.Repo.Create(ctx, u)
	if err != nil {
		log.Fatal("Something wrong with repo")
	}

	return value
}
