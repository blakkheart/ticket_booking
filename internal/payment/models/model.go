package models

import "ticket-booking/internal/money"

type PaymentStatus string

const (
	Pending   PaymentStatus = "pending"
	Succeeded PaymentStatus = "succeeded"
	Failed    PaymentStatus = "failed"
)

type Payment struct {
	ID          int64
	BookingID   int64
	Amount      money.Money
	Status      PaymentStatus
	Provider    string
	OperationID string
}
