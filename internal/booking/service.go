package booking

import (
	"context"

)

type Service interface {
	Get(id int64) *Booking
	GetMany(filters any) []*Booking
	Create(ctx context.Context, user *Booking) *Booking
}

func NewService(repo Repository) Service {
	return &bookingService{Repo: repo}
}

type bookingService struct {
	Repo Repository
}

func (service *bookingService) Get(id int64) *Booking {
	return nil
}

func (service *bookingService) GetMany(filters any) []*Booking {
	return nil
}

func (service *bookingService) Create(ctx context.Context, user *Booking) *Booking {
	return nil
}
