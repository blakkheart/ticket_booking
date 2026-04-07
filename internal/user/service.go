package user

import (
	"context"
	"fmt"
	"log/slog"
	"ticket-booking/internal/repository"
	"ticket-booking/internal/user/models"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Get(id int64) (*models.User, error)
	GetMany(filters any) []*models.User
	Create(ctx context.Context, user *models.CreateUserRequest) (*models.User, error)
	GetUserAuthByEmail(ctx context.Context, email string) (*models.UserAuth, error)
}

func NewService(repo repository.UserRepository, logger *slog.Logger) Service {
	return &userService{Repo: repo, logger: logger}
}

type userService struct {
	Repo   repository.UserRepository
	logger *slog.Logger
}

func (s *userService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *userService) Get(id int64) (*models.User, error) {
	value, err := s.Repo.Get(id)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (s *userService) GetMany(filters any) []*models.User {
	return nil
}

func (s *userService) GetUser() models.User {
	return models.User{
		ID:    1,
		Name:  "Name",
		Email: "Email",
	}
}

func (s *userService) Create(ctx context.Context, user *models.CreateUserRequest) (*models.User, error) {
	hashedPassword, err := s.hashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	u := &models.CreateUserRequest{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	}

	value, err := s.Repo.Create(ctx, u)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	return value, nil
}

func (s *userService) GetUserAuthByEmail(ctx context.Context, email string) (*models.UserAuth, error) {
	value, err := s.Repo.GetByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return value, nil
}
