package ticket

import "context"

type Repository interface {
	Create(ctx context.Context, dto *TicketTypeIn) (*TicketType, error)
	Delete(id int64) error
	Get(id int64) (*TicketType, error)
	GetMany(filter any) ([]*TicketType, error)
}
