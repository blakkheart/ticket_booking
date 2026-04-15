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
	BookingID  int64
	Amount     money.Money
	Currency   string
	Status     PaymentStatus
	Provider   string
	ExternalID string // айди в системе провайдера
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
