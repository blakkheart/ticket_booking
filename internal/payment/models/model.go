package models

import (
	"ticket-booking/internal/money"
	"time"
)

type PaymentStatus string

const (
	PaymentStatusCreated   PaymentStatus = "created"
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusSucceeded PaymentStatus = "succeeded"
	PaymentStatusFailed    PaymentStatus = "failed"
)

type Payment struct {
	ID         string
	OrderID    string
	Amount     money.Money
	Currency   string
	Status     PaymentStatus
	Provider   string
	ExternalID string // айди в системе провайдера
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CreatePaymentRequest struct {
	PaymentID   string
	Amount      money.Money
	Currency    string
	Description string

	// куда редиректить пользователя после оплаты
	SuccessURL string
	FailURL    string
}

type CreatePaymentResponse struct {
	PaymentID  string
	ExternalID string

	// если redirect платеж например страйп
	PaymentURL string

	// если SDK
	ClientToken string
}

type CallbackData struct {
	PaymentID  string
	ExternalID string

	Status   PaymentStatus
	Amount   money.Money
	Currency string
}
