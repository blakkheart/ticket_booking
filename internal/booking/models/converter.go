package models

import (
	"errors"
	"log/slog"

	"github.com/google/uuid"
)

func ConvertRequestToParams(req *CreateBookingRequest) (*CreateBookingParams, error) {
	eventUUID, err := uuid.Parse(req.EventID)
	if err != nil {
		slog.Error("Event id is invalid", "error", err)
		return nil, errors.New("Event id is invalid")
	}

	itemsParams := make([]CreateBookingItemParams, 0, len(req.Items))
	for _, item := range req.Items {
		tickeTypeUUID, err := uuid.Parse(item.TicketTypeID)
		if err != nil {
			slog.Error("TicketType id is invalid", "error", err)
			return nil, errors.New("TicketType id is invalid")
		}
		itemsParams = append(itemsParams, CreateBookingItemParams{
			TicketTypeID: tickeTypeUUID,
			Quantity:     item.Quantity,
		})
	}

	return &CreateBookingParams{
		EventID: eventUUID,
		Items:   itemsParams,
	}, nil
}
