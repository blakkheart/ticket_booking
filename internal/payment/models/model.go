package models

import (
	"ticket-booking/internal/money"
	"time"

	"github.com/google/uuid"
)

type PaymentStatus string

const (
	PaymentStatusCreated   PaymentStatus = "created"
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusSucceeded PaymentStatus = "succeeded"
	PaymentStatusFailed    PaymentStatus = "failed"
)

type Payment struct {
	ID              uuid.UUID
	PaymentIntentID uuid.UUID
	BookingID       uuid.UUID
	Amount          money.Money
	Currency        string
	Status          PaymentStatus
	Provider        string
	ExternalID      string // айди в системе провайдера
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type CreatePaymentParams struct {
	ID          uuid.UUID
	Amount      money.Money
	Currency    string
	Description string

	SuccessURL string
	FailURL    string
}
