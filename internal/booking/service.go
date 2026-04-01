package booking

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	bookingitems "ticket-booking/internal/booking_items"
	"ticket-booking/internal/event"
	"ticket-booking/internal/money"
	"ticket-booking/internal/ticket"
	"time"
)

type CreateBookingRequest struct {
	EventID int                    `json:"event_id"`
	Items   []CreateBookingItemDTO `json:"items"`
}

type CreateBookingItemDTO struct {
	TicketTypeID int `json:"ticket_type_id"`
	Quantity     int `json:"quantity"`
}

type Service interface {
	Get(id int64) *Booking
	GetMany(filters any) []*Booking
	Create(ctx context.Context, b *BookingIn) (*Booking, error)
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

func (service *bookingService) Create(ctx context.Context, b *BookingIn) (*Booking, error) {

	request := `{
	  "event_id": 123,
	  "items": [
	    {
	      "ticket_type_id": 1,
	      "quantity": 2
	    }
	  ]
	}`

	var breq CreateBookingRequest
	errorJson := json.Unmarshal([]byte(request), &breq)
	if errorJson != nil {
		return nil, errorJson
	}

	accountID := 1
	ticketID := 1

	expiresAfter := time.Now().Add(15 * time.Minute)

	// pending     — создана, ждет оплаты
	// confirmed   — оплачена
	// cancelled   — отменена
	// expired     — не оплатил вовремя

	bookingIn := BookingIn{
		AccountID: int64(accountID),
		Status:    "status",
		ExpiresAt: &expiresAfter,
		PaidAt:    nil,
	}
	bookingCreated, err := service.bookingRepo.Create(ctx, &bookingIn)

	if err != nil {
		return nil, errors.New("Cannot create booking")
	}

	price, _ := money.NewMoney("1")

	bookingItemIn := bookingitems.BookingItemIn{
		BookingID:      bookingCreated.ID,
		TicketTypeID:   int64(ticketID),
		Quantity:       1,
		PriceAtBooking: *price,
	}

	_, errС := service.bookingItemsRepo.Create(ctx, &bookingItemIn)
	if errС != nil {
		if errors.Is(errС, BookingItemsAlreadyExists) {
			return nil, errС
		}
		return nil, errors.New("Cannot create bookingItem")
	}

	// 	{    response
	//   "booking_id": 555,
	//   "status": "pending",
	//   "total_price": 260
	// }

	return bookingCreated, nil
}
