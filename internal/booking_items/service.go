package bookingitems

import (
	"context"
	"log/slog"
	"ticket-booking/internal/booking_items/models"
	"ticket-booking/internal/repository"
)

type Service interface {
	Get(id int64) *models.BookingItem
	GetMany(filters any) []*models.BookingItem
	Create(ctx context.Context, b *models.BookingItemIn) *models.BookingItem
}

func NewService(repo repository.BookingItemsRepository, logger *slog.Logger) Service {
	return &bookingItemsService{Repo: repo, logger: logger}
}

type bookingItemsService struct {
	Repo   repository.BookingItemsRepository
	logger *slog.Logger
}

func (service *bookingItemsService) Get(id int64) *models.BookingItem {
	return nil
}

func (service *bookingItemsService) GetMany(filters any) []*models.BookingItem {
	return nil
}

func (service *bookingItemsService) Create(ctx context.Context, b *models.BookingItemIn) *models.BookingItem {
	return nil
}
