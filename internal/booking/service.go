package booking

import (
	"context"
	"errors"
	"log/slog"
	bookingitems "ticket-booking/internal/booking_items"
	"ticket-booking/internal/event"
	"ticket-booking/internal/ticket"
	"time"
)

type Service interface {
	Get(id int64) *Booking
	GetMany(filters any) []*Booking
	Create(ctx context.Context, b *CreateBookingRequest) (*Booking, error)
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

func (service *bookingService) Get(id int64) *Booking {
	return nil
}

func (service *bookingService) GetMany(filters any) []*Booking {
	return nil
}

func (service *bookingService) Create(ctx context.Context, b *CreateBookingRequest) (*Booking, error) {

	accountID := int64(1)

	expiresAfter := time.Now().Add(15 * time.Minute)

	bookingIn := BookingIn{
		AccountID: accountID,
		Status:    Pending,
		ExpiresAt: &expiresAfter,
		PaidAt:    nil,
	}
	bookingCreated, err := service.bookingRepo.Create(ctx, &bookingIn)

	if err != nil {
		return nil, errors.New("Cannot create booking")
	}

	for _, bItem := range b.Items {

		ticketType, errTicket := service.ticketRepo.Get(ctx, bItem.TicketTypeID)

		if errTicket != nil {
			return nil, errors.New("TicketType dosent exist")
		}
		if ticketType.EventID != b.EventID {
			return nil, errors.New("TicketType dosent match eventID")
		}

		bookingItemIn := bookingitems.BookingItemIn{
			BookingID:      bookingCreated.ID,
			TicketTypeID:   bItem.TicketTypeID,
			Quantity:       bItem.Quantity,
			PriceAtBooking: ticketType.Price,
		}

		_, errС := service.bookingItemsRepo.Create(ctx, &bookingItemIn)
		if errС != nil {
			if errors.Is(errС, BookingItemsAlreadyExists) {
				return nil, errС
			}
			return nil, errors.New("Cannot create bookingItem")
		}
	}

	return bookingCreated, nil
}
