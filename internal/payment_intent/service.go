package paymentintent

import (
	"context"
	"log/slog"
	bookingmodel "ticket-booking/internal/booking/models"
	"ticket-booking/internal/payment_intent/models"
	"ticket-booking/internal/repository"
	"ticket-booking/internal/uow"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	Create(
		ctx context.Context,
		bookingID int64,
		userID int64,
	) (*models.PaymentIntent, error)
}

func NewService(
	repo repository.PaymentIntentRepository,
	uowManager uow.UnitOfWorkManager,
	logger *slog.Logger,
) Service {
	return &paymentIntentService{
		Repo:       repo,
		uowManager: uowManager,
		logger:     logger,
	}
}

type paymentIntentService struct {
	Repo       repository.PaymentIntentRepository
	logger     *slog.Logger
	uowManager uow.UnitOfWorkManager
}

func (s *paymentIntentService) Create(
	ctx context.Context,
	bookingID int64,
	userID int64,
	// idempotencyKey string,
) (*models.PaymentIntent, error) {

	var result *models.PaymentIntent

	err := s.uowManager.Do(ctx, func(uow uow.UnitOfWork) error {
		booking, err := uow.BookingRepo().Get(ctx, bookingID)
		if err != nil {
			return err
		}

		if booking.Status != bookingmodel.StatusPending {
			return ErrInvalidBookingState
		}

		if time.Now().After(*booking.ExpiresAt) {
			return ErrBookingExpired
		}

		existing, _ := uow.PaymentIntentRepo().GetByBookingID(ctx, bookingID)
		if existing != nil {
			result = existing
			return nil
		}

		intent := &models.PaymentIntent{
			ID:        uuid.NewString(),
			BookingID: booking.ID,
			Amount:    *booking.TotalPrice,
			Currency:  "RUB",
			Status:    models.IntentRequiresPayment,
			ExpiresAt: time.Now().Add(15 * time.Minute),
		}

		if err := uow.PaymentIntentRepo().Create(ctx, intent); err != nil {
			return err
		}

		booking.Status = bookingmodel.StatusAwaitingPayment
		if err := uow.BookingRepo().Update(ctx, booking); err != nil {
			return err
		}

		result = intent
		return nil
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}
