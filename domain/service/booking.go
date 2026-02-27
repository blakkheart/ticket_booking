package service

import (
	"context"

	model "ticket-booking/models/domain/booking"
	model_interface "ticket-booking/models/domain/interface"
)

type BookingService interface {
	Get(id int64) *model.Booking
	GetMany(filters any) []*model.Booking
	Create(ctx context.Context, user *model.Booking) *model.Booking
}

func NewB(repo model_interface.IBookingRepository) BookingService {
	return &bookingService{Repo: repo}
}

type bookingService struct {
	Repo model_interface.IBookingRepository
}

func (service *bookingService) Get(id int64) *model.Booking {
	return nil
}

func (service *bookingService) GetMany(filters any) []*model.Booking {
	return nil
}

func (service *bookingService) Create(ctx context.Context, user *model.Booking) *model.Booking {
	return nil
}
