package ticket

import "context"

type Repository interface {
	Create(ctx context.Context, dto *Ticket) (*Ticket, error)
	Delete(id int64) error
	Get(id int64) (*Ticket, error)
	GetMany(filter any) ([]*Ticket, error)
}
