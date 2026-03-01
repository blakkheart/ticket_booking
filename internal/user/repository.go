package user

import "context"

type Repository interface {
	Create(ctx context.Context, dto *AccountIn) (*Account, error)
	Delete(id int64) error
	Get(id int64) (*Account, error)
	GetMany(filter any) ([]*Account, error)
}
