package service

import (
	"log"

	"golang.org/x/crypto/bcrypt"

	model "ticket-booking/models/domain/account"
	model_interface "ticket-booking/models/domain/interface"
	"ticket-booking/repository"
)

type UserService struct {
	Repo model_interface.IAccountRepository
}

func (service *UserService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (service *UserService) verifyPassword(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}

func (service *UserService) Get(id int64) *model.Account {
	value, err := service.Repo.Get(id)
	if err != nil {
		log.Fatal("Something wrong with repo")
	}
	return value
}

func (service *UserService) GetMany(filters any) {
}

func (service *UserService) GetUser() repository.Account {
	return repository.Account{
		ID:    1,
		Name:  "Name",
		Email: "Email",
	}
}

func (service *UserService) Create(user *model.AccountIn) *model.Account {

	hashedPassword, err := service.hashPassword(user.Password)

	if err != nil {
		log.Fatal("Cannot hash password")
	}

	user.Password = hashedPassword

	u := &model.AccountIn{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	value, err := service.Repo.Create(u)
	if err != nil {
		log.Fatal("Something wrong with repo")
	}

	return value
}

func (service *UserService) authorization(user repository.Account) bool {
	return true
}
