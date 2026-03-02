package user

import "context"

type Repository interface {
	Create(ctx context.Context, dto *CreateUserRequest) (*User, error)
	Delete(id int64) error
	Get(id int64) (*User, error)
	GetMany(filter any) ([]*User, error)
}
