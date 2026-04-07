package booking

import (
	"context"
	"errors"
	"log/slog"
	bookingitems "ticket-booking/internal/booking_items"
	"ticket-booking/internal/event"
	"ticket-booking/internal/money"
	"ticket-booking/internal/ticket"
	"time"
)

type Service interface {
	Get(id int64) (*Booking, error)
	GetMany(filters any) []*Booking
	Create(ctx context.Context, b *CreateBookingRequest, userID int64) (*Booking, error)
}

func NewService(
	bookingRepo Repository,
	eventRepo event.Repository,
	bookingItemsRepo bookingitems.Repository,
	ticketRepo ticket.Repository,
	logger *slog.Logger,
) Service {
	return &bookingService{
		bookingRepo:      bookingRepo,
		eventRepo:        eventRepo,
		bookingItemsRepo: bookingItemsRepo,
		ticketRepo:       ticketRepo,
		logger:           logger,
	}
}

type bookingService struct {
	bookingRepo      Repository
	eventRepo        event.Repository
	bookingItemsRepo bookingitems.Repository
	ticketRepo       ticket.Repository
	logger           *slog.Logger
}

func (service *bookingService) Get(id int64) (*Booking, error) {
	booking, err := service.bookingRepo.Get(id)
	return booking, err
}

func (service *bookingService) GetMany(filters any) []*Booking {
	return nil
}

func (service *bookingService) Create(ctx context.Context, b *CreateBookingRequest, userID int64) (*Booking, error) {

	expiresAfter := time.Now().Add(15 * time.Minute)

	bookingIn := BookingIn{
		AccountID: userID,
		Status:    Pending,
		ExpiresAt: &expiresAfter,
		PaidAt:    nil,
	}
	bookingCreated, err := service.bookingRepo.Create(ctx, &bookingIn)

	if err != nil {
		return nil, errors.New("Cannot create booking")
	}

	totalPrice, errMoney := money.NewMoney("0")
	if errMoney != nil {
		return nil, errMoney
	}

	for _, bItem := range b.Items {

		ticketType, errTicket := service.ticketRepo.Get(ctx, bItem.TicketTypeID)

		if errTicket != nil {
			return nil, errors.New("TicketType dosent exist")
		}
		if ticketType.EventID != b.EventID {
			return nil, errors.New("TicketType dosent match eventID")
		}

		if ticketType.AvailableQuantity < bItem.Quantity {
			return nil, errors.New("Avaliable ticket quantity is too low")
		}

		totalPrice = totalPrice.Add(&ticketType.Price)

		bookingItemIn := bookingitems.BookingItemIn{
			BookingID:      bookingCreated.ID,
			TicketTypeID:   bItem.TicketTypeID,
			Quantity:       bItem.Quantity,
			PriceAtBooking: ticketType.Price,
		}

		service.ticketRepo.UpdateTicketQuantityByID(ctx, bItem.TicketTypeID, ticketType.AvailableQuantity-bItem.Quantity)

		_, errC := service.bookingItemsRepo.Create(ctx, &bookingItemIn)
		if errC != nil {
			if errors.Is(errC, BookingItemsAlreadyExists) {
				return nil, errC
			}
			return nil, errors.New("Cannot create bookingItem")
		}
	}

	bookingCreated.TotalPrice = totalPrice

	return bookingCreated, nil
}
