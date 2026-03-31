package booking

import (
	"context"
	"log/slog"
)

type Service interface {
	Get(id int64) *Booking
	GetMany(filters any) []*Booking
	Create(ctx context.Context, b *BookingIn) *Booking
}

func NewService(repo Repository, logger *slog.Logger) Service {
	return &bookingService{Repo: repo, logger: logger}
}

type bookingService struct {
	Repo   Repository
	logger *slog.Logger
}

func (service *bookingService) Get(id int64) *Booking {
	return nil
}

func (service *bookingService) GetMany(filters any) []*Booking {
	return nil
}

func (service *bookingService) Create(ctx context.Context, b *BookingIn) *Booking {
	return nil
}
