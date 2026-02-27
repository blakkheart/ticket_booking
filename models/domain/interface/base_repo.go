package model_interface

import "context"

type IBaseRepository[T any, CreateDTO any] interface {
	Create(ctx context.Context, dto *CreateDTO) (*T, error)
	Delete(id int64) error
	Get(id int64) (*T, error)
	GetMany(filter any) ([]*T, error)
}
