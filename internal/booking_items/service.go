package bookingitems

import (
	"context"
	"log/slog"
	"ticket-booking/internal/booking_items/models"
	"ticket-booking/internal/repository"

	"github.com/google/uuid"
)

type Service interface {
	Get(id uuid.UUID) *models.BookingItem
	GetMany(filters any) []*models.BookingItem
	Create(
		ctx context.Context,
		b *models.BookingItemIn,
	) *models.BookingItem
}

func NewService(
	repo repository.BookingItemsRepository,
	logger *slog.Logger,
) Service {
	return &bookingItemsService{Repo: repo, logger: logger}
}

type bookingItemsService struct {
	Repo   repository.BookingItemsRepository
	logger *slog.Logger
}

func (service *bookingItemsService) Get(id uuid.UUID) *models.BookingItem {
	return nil
}

func (service *bookingItemsService) GetMany(filters any) []*models.BookingItem {
	return nil
}

func (service *bookingItemsService) Create(
	ctx context.Context,
	b *models.BookingItemIn,
) *models.BookingItem {
	return nil
}
