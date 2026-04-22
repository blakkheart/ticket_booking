package booking

import (
	"context"
	"errors"
	"log/slog"
	"ticket-booking/internal/booking/models"
	biModels "ticket-booking/internal/booking_items/models"
	"ticket-booking/internal/money"
	"ticket-booking/internal/repository"
	"ticket-booking/internal/ticket"
	"ticket-booking/internal/uow"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	Get(id uuid.UUID) (*models.Booking, error)
	GetMany(filters any) []*models.Booking
	Create(
		ctx context.Context,
		b *models.CreateBookingParams,
		userID uuid.UUID,
	) (*models.Booking, error)
}

func NewService(
	bookingRepo repository.BookingRepository,
	eventRepo repository.EventRepository,
	bookingItemsRepo repository.BookingItemsRepository,
	ticketService ticket.Service,
	uowManager uow.UnitOfWorkManager,
	logger *slog.Logger,
) Service {
	return &bookingService{
		bookingRepo:      bookingRepo,
		eventRepo:        eventRepo,
		bookingItemsRepo: bookingItemsRepo,
		ticketService:    ticketService,
		uowManager:       uowManager,
		logger:           logger,
	}
}

type bookingService struct {
	bookingRepo      repository.BookingRepository
	eventRepo        repository.EventRepository
	bookingItemsRepo repository.BookingItemsRepository
	ticketService    ticket.Service
	uowManager       uow.UnitOfWorkManager

	logger *slog.Logger
}

func (service *bookingService) Get(id uuid.UUID) (*models.Booking, error) {
	booking, err := service.bookingRepo.Get(context.TODO(), id)
	return booking, err
}

func (service *bookingService) GetMany(filters any) []*models.Booking {
	return nil
}

func (service *bookingService) Create(
	ctx context.Context,
	b *models.CreateBookingParams,
	userID uuid.UUID,
) (res *models.Booking, err error) {

	var result *models.Booking

	err = service.uowManager.Do(ctx, func(uow uow.UnitOfWork) error {
		bookingRepo := uow.BookingRepo()
		bookingItemsRepo := uow.BookingItemsRepo()

		expiresAfter := time.Now().Add(15 * time.Minute)

		bookingIn := models.BookingIn{
			ID:        uuid.New(),
			AccountID: userID,
			Status:    models.StatusPending,
			ExpiresAt: &expiresAfter,
			PaidAt:    nil,
		}
		res, err = bookingRepo.Create(ctx, &bookingIn)

		if err != nil {
			return errors.New("Cannot create booking")
		}

		totalPrice, err := money.NewMoney("0")
		if err != nil {
			return err
		}

		for _, bItem := range b.Items {

			ticketType, err := service.ticketService.Reserve(
				ctx,
				uow,
				bItem.TicketTypeID,
				bItem.Quantity,
				b.EventID,
			)
			if err != nil {
				return err
			}

			totalPrice = totalPrice.Add(&ticketType.Price)

			bookingItemIn := biModels.BookingItemIn{
				BookingID:      res.ID,
				TicketTypeID:   bItem.TicketTypeID,
				Quantity:       bItem.Quantity,
				PriceAtBooking: ticketType.Price,
			}

			_, err = bookingItemsRepo.Create(ctx, &bookingItemIn)
			if err != nil {
				if errors.Is(err, BookingItemsAlreadyExists) {
					return err
				}
				return errors.New("Cannot create bookingItem")
			}
		}

		res.TotalPrice = totalPrice

		result = res
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
