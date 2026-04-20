package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID
	Title       string
	Description *string
	Location    string
	StartsAt    time.Time
	EndsAt      *time.Time
	Status      string
}

type EventIn struct {
	ID          uuid.UUID
	Title       string
	Description *string
	Location    string
	StartsAt    time.Time
	EndsAt      *time.Time
	Status      string
}
