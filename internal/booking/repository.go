package booking

import "context"

type Repository interface {
	Create(ctx context.Context, dto *Booking) (*Booking, error)
	Delete(id int64) error
	Get(id int64) (*Booking, error)
	GetMany(filter any) ([]*Booking, error)
}
