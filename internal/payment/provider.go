package payment

import (
	"context"
	"errors"
	"net/http"
	"ticket-booking/internal/payment/models"
)

type PaymentProvider interface {
	Name() string
	CreatePayment(ctx context.Context, req models.CreatePaymentRequest) (*models.CreatePaymentResponse, error)
	VerifySignature(r *http.Request) error
	ParseCallback(r *http.Request) (*models.CallbackData, error)
}

type MockProvider struct {
	secret string
}

func (p *MockProvider) ParseCallback(r *http.Request) (*models.CallbackData, error) {
	return &models.CallbackData{}, nil
}

func (p *MockProvider) VerifySignature(r *http.Request) error {
	sig := r.Header.Get("X-Signature")

	expected, err := generateHMAC(r.Body, p.secret)

	if err != nil {
		return err
	}

	if sig != expected {
		return errors.New("invalid signature")
	}

	return nil
}

func (p *MockProvider) CreatePayment(ctx context.Context, req models.CreatePaymentRequest) (*models.CreatePaymentResponse, error) {
	return &models.CreatePaymentResponse{
		PaymentURL: "http://localhost:8080/mock/pay",
	}, nil
}

func (p *MockProvider) Name() string {
	return "mock"
}

func NewMockProvider(secret string) *MockProvider {
	return &MockProvider{
		secret: secret,
	}
}
