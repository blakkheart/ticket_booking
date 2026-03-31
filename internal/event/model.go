package event

import (
	"time"
)

type Event struct {
	ID          int64
	Title       string
	Description *string
	Location    string
	StartsAt    time.Time
	EndsAt      *time.Time
	Status      string
}

type EventIn struct {
	Title       string
	Description *string
	Location    string
	StartsAt    time.Time
	EndsAt      *time.Time
	Status      string
}
