package model_interface

// type BaseRepository interface {
// 	Create(user repository.Account) (repository.Account, error)
// 	Delete(id int) (int, error)
// 	Get(id int) (int, error)
// 	GetMany(any)
// 	CreateEvent(repository.CreateEventParams) (repository.Event, error)
// 	GetEvent(id int64) (repository.Event, error)
// }

type IBaseRepository[T any, CreateDTO any] interface {
	Create(dto CreateDTO) (T, error)
	Delete(id int64) error
	Get(id int64) (T, error)
	GetMany(filter any) ([]T, error)
}
