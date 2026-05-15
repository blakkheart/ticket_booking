package repository

import (
	"context"
	"ticket-booking/internal/user/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(
		ctx context.Context,
		dto *models.CreateUserParams,
	) (*models.User, error)
	Delete(id uuid.UUID) error
	Get(id uuid.UUID) (*models.User, error)
	GetMany(filter any) ([]*models.User, error)
	GetByEmail(
		ctx context.Context,
		email string,
	) (*models.UserAuth, error)
}
