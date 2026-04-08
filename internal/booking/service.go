package booking

import (
	"context"
	"errors"
	"log/slog"
	"ticket-booking/internal/booking/models"
	biModels "ticket-booking/internal/booking_items/models"
	"ticket-booking/internal/money"
	"ticket-booking/internal/repository"
	"ticket-booking/internal/uow"
	"time"
)

type Service interface {
	Get(id int64) (*models.Booking, error)
	GetMany(filters any) []*models.Booking
	Create(ctx context.Context, b *models.CreateBookingRequest, userID int64) (*models.Booking, error)
}

func NewService(
	bookingRepo repository.BookingRepository,
	eventRepo repository.EventRepository,
	bookingItemsRepo repository.BookingItemsRepository,
	ticketRepo repository.TicketRepository,
	uowManager uow.UnitOfWorkManager,
	logger *slog.Logger,
) Service {
	return &bookingService{
		bookingRepo:      bookingRepo,
		eventRepo:        eventRepo,
		bookingItemsRepo: bookingItemsRepo,
		ticketRepo:       ticketRepo,
		uowManager:       uowManager,
		logger:           logger,
	}
}

type bookingService struct {
	bookingRepo      repository.BookingRepository
	eventRepo        repository.EventRepository
	bookingItemsRepo repository.BookingItemsRepository
	ticketRepo       repository.TicketRepository
	uowManager       uow.UnitOfWorkManager

	logger *slog.Logger
}

func (service *bookingService) Get(id int64) (*models.Booking, error) {
	booking, err := service.bookingRepo.Get(id)
	return booking, err
}

func (service *bookingService) GetMany(filters any) []*models.Booking {
	return nil
}

func (service *bookingService) Create(ctx context.Context, b *models.CreateBookingRequest, userID int64) (res *models.Booking, err error) {

	uowM, err := service.uowManager.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			service.logger.Debug("Transaction rollback")
			_ = uowM.Rollback(ctx)
			return
		}
		service.logger.Debug("Transaction commit")
		err = uowM.Commit(ctx)
	}()

	bookingRepo := uowM.BookingRepo()
	ticketRepo := uowM.TicketRepo()
	bookingItemsRepo := uowM.BookingItemsRepo()

	expiresAfter := time.Now().Add(15 * time.Minute)

	bookingIn := models.BookingIn{
		AccountID: userID,
		Status:    models.Pending,
		ExpiresAt: &expiresAfter,
		PaidAt:    nil,
	}
	res, err = bookingRepo.Create(ctx, &bookingIn)

	if err != nil {
		return nil, errors.New("Cannot create booking")
	}

	totalPrice, err := money.NewMoney("0")
	if err != nil {
		return nil, err
	}

	for _, bItem := range b.Items {

		ticketType, err := ticketRepo.Get(ctx, bItem.TicketTypeID)

		if err != nil {
			service.logger.Error("ticket type doesn't exist: %w", err)
			return nil, errors.New("TicketType dosent exist")
		}
		if ticketType.EventID != b.EventID {
			return nil, errors.New("TicketType dosent match eventID")
		}

		if ticketType.AvailableQuantity < bItem.Quantity {
			return nil, errors.New("Avaliable ticket quantity is too low")
		}

		totalPrice = totalPrice.Add(&ticketType.Price)

		bookingItemIn := biModels.BookingItemIn{
			BookingID:      res.ID,
			TicketTypeID:   bItem.TicketTypeID,
			Quantity:       bItem.Quantity,
			PriceAtBooking: ticketType.Price,
		}

		ticketRepo.UpdateTicketQuantityByID(ctx, bItem.TicketTypeID, ticketType.AvailableQuantity-bItem.Quantity)

		_, err = bookingItemsRepo.Create(ctx, &bookingItemIn)
		if err != nil {
			if errors.Is(err, BookingItemsAlreadyExists) {
				return nil, err
			}
			return nil, errors.New("Cannot create bookingItem")
		}
	}

	res.TotalPrice = totalPrice

	return res, nil
}
