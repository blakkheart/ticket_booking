package service

import (
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"

	model "ticket-booking/models/domain/account"
	model_interface "ticket-booking/models/domain/interface"
)

type UserService interface {
	Get(id int64) *model.Account
	GetMany(filters any) []*model.Account
	Create(ctx context.Context, user *model.AccountIn) *model.Account
}

func New(repo model_interface.IAccountRepository) UserService {
	return &userService{Repo: repo}
}

type userService struct {
	Repo model_interface.IAccountRepository
}

func (service *userService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (service *userService) verifyPassword(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}

func (service *userService) Get(id int64) *model.Account {
	value, err := service.Repo.Get(id)
	if err != nil {
		log.Fatal("Something wrong with repo")
	}
	return value
}

func (service *userService) GetMany(filters any) []*model.Account {
	return nil
}

func (service *userService) GetUser() model.Account {
	return model.Account{
		ID:    1,
		Name:  "Name",
		Email: "Email",
	}
}

func (service *userService) Create(ctx context.Context, user *model.AccountIn) *model.Account {
	hashedPassword, err := service.hashPassword(user.Password)

	if err != nil {
		log.Fatal("Cannot hash password")
	}

	u := &model.AccountIn{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	}

	value, err := service.Repo.Create(ctx, u)
	if err != nil {
		log.Fatal("Something wrong with repo")
	}

	return value
}

func (service *userService) authorization(user *model.Account) bool {
	return true
}
