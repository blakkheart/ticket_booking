package event

import "context"

type Repository interface {
	Create(ctx context.Context, dto *EventIn) (*Event, error)
	Delete(id int64) error
	Get(id int64) (*Event, error)
	GetMany(filter any) ([]*Event, error)
}
