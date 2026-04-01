package bookingitems

import "context"

type Repository interface {
	Create(ctx context.Context, dto *BookingItemIn) (*BookingItem, error)
	Delete(id int64) error
	Get(id int64) (*BookingItem, error)
	GetMany(filter any) ([]*BookingItem, error)
}
