package bookingitems

import (
	"context"
	"log/slog"
)

type Service interface {
	Get(id int64) *BookingItem
	GetMany(filters any) []*BookingItem
	Create(ctx context.Context, b *BookingItemIn) *BookingItem
}

func NewService(repo Repository, logger *slog.Logger) Service {
	return &bookingItemsService{Repo: repo, logger: logger}
}

type bookingItemsService struct {
	Repo   Repository
	logger *slog.Logger
}

func (service *bookingItemsService) Get(id int64) *BookingItem {
	return nil
}

func (service *bookingItemsService) GetMany(filters any) []*BookingItem {
	return nil
}

func (service *bookingItemsService) Create(ctx context.Context, b *BookingItemIn) *BookingItem {
	return nil
}
