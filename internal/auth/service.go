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
	user := a.userServise.GetByEmail(email)
	a.checkPassword(password, "1")
	return a.generateToken(user.ID, string(user.Role))
}

func (a *authService) checkPassword(password string, hashedPassword string) bool {
	return true
}

func (a *authService) generateToken(userID int64, role string) (string, error) {
	return a.jwt.Generate(userID, role)
}
