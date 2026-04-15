package models

import "ticket-booking/internal/money"

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
