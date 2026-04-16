package payment

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"ticket-booking/internal/config"
	"ticket-booking/internal/payment/models"
	"ticket-booking/internal/repository"
	"time"

	"github.com/google/uuid"
)

func generateHMAC(body io.ReadCloser, secret string) (string, error) {
	defer body.Close()

	data, err := io.ReadAll(body)
	if err != nil {
		return "", err
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(data)

	return hex.EncodeToString(mac.Sum(nil)), nil
}

type Service struct {
	repo repository.PaymentRepository
}

func (s *Service) HandleCallback(ctx context.Context, provider PaymentProvider, r *http.Request) error {
	if err := provider.VerifySignature(r); err != nil {
		return err
	}

	data, err := provider.ParseCallback(r)
	if err != nil {
		return err
	}

	payment, err := s.repo.GetByID(ctx, data.PaymentID)
	if err != nil {
		return err
	}

	if payment.Status == models.PaymentStatusSucceeded {
		return nil
	}

	payment.Status = data.Status
	return s.repo.Update(ctx, payment)
}

func (s *Service) CreatePayment(
	ctx context.Context,
	provider PaymentProvider,
	req models.CreatePaymentRequest,
) (*models.CreatePaymentResponse, error) {

	payment := &models.Payment{
		ID:        uuid.NewString(),
		PaymentID: req.PaymentID,
		Amount:    req.Amount,
		Currency:  req.Currency,
		Status:    models.PaymentStatusPending,
		Provider:  provider.Name(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, payment); err != nil {
		return nil, err
	}

	resp, err := provider.CreatePayment(ctx, models.CreatePaymentRequest{
		PaymentID:   payment.ID,
		Amount:      payment.Amount,
		Currency:    payment.Currency,
		Description: req.Description,
		SuccessURL:  req.SuccessURL,
		FailURL:     req.FailURL,
	})
	if err != nil {
		return nil, err
	}

	payment.ExternalID = resp.ExternalID
	_ = s.repo.Update(ctx, payment)

	return &models.CreatePaymentResponse{
		PaymentID:   payment.ID,
		PaymentURL:  resp.PaymentURL,
		ClientToken: resp.ClientToken,
	}, nil
}

func NewProvider(cfg config.PaymentConfig) PaymentProvider {
	switch cfg.Provider {
	case "mock":
		return NewMockProvider(
			cfg.Secret,
		)
		// case "stripe":
		// 	return NewStripeProvider()
		// }
	}
	return nil
}

// POST /payments  создать платеж
// GET  /payments/{id}  статус
// POST /payments/callback/{provider}
