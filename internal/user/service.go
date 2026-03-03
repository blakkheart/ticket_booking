package user

import (
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Get(id int64) *User
	GetMany(filters any) []*User
	Create(ctx context.Context, user *CreateUserRequest) (*User, error)
	GetByEmail(email string, password string) *User
}

func NewService(repo Repository) Service {
	return &userService{Repo: repo}
}

type userService struct {
	Repo Repository
}

func (service *userService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (service *userService) verifyPassword(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}

func (service *userService) Get(id int64) *User {
	value, err := service.Repo.Get(id)
	if err != nil {
		log.Fatal("Something wrong with repo")
	}
	return value
}

func (service *userService) GetMany(filters any) []*User {
	return nil
}

func (service *userService) GetUser() User {
	return User{
		ID:    1,
		Name:  "Name",
		Email: "Email",
	}
}

func (service *userService) Create(ctx context.Context, user *CreateUserRequest) (*User, error) {
	// hashedPassword, err := service.hashPassword(user.Password)

	// if err != nil {
	// 	log.Fatal("Cannot hash password")
	// }

	// u := &CreateUserRequest{
	// 	Name:     user.Name,
	// 	Email:    user.Email,
	// 	Password: hashedPassword,
	// }

	// value, err := service.Repo.Create(ctx, u)
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// 	log.Fatal("Something wrong with repo")
	// }

	// return value, nil
	return nil, ErrEmailAlreadyUsed
}

func (service *userService) authorization(user *User) bool {
	return true
}

func (s *userService) GetByEmail(email string, password string) *User {
	return nil
}
