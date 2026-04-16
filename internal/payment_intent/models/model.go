package models

import (
	"ticket-booking/internal/money"
	"time"
)

type IntentStatus string

const (
	IntentRequiresPayment IntentStatus = "requires_payment"
	IntentProcessing      IntentStatus = "processing"
	IntentSucceeded       IntentStatus = "succeeded"
	IntentFailed          IntentStatus = "failed"
	IntentCancelled       IntentStatus = "cancelled"
)

type PaymentIntent struct {
	ID        int64
	BookingID int64
	Amount    money.Money
	Currency  string
	Status    IntentStatus
	ExpiresAt time.Time
	CreatedAt time.Time
}
