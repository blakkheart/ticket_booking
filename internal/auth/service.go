package auth

import (
	"context"
	"ticket-booking/internal/user"
)

type Service interface {
	Login(ctx context.Context, email string, password string) (string, error)
}

func NewService(jwt *JWTManager, userService user.Service) Service {
	return &authService{
		jwt:         jwt,
		userServise: userService,
	}
}

type authService struct {
	jwt         *JWTManager
	userServise user.Service
}

func (a *authService) Login(ctx context.Context, email string, password string) (string, error) {
	user := a.userServise.GetByEmail(email, password)
	a.checkPassword(password)
	return a.generatToken(user.ID, string(user.Role))
}

func (a *authService) checkPassword(password string) bool {
	return true
}

func (a *authService) generatToken(userID int64, role string) (string, error) {
	return a.jwt.Generate(userID, role)
}
