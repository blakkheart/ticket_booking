package auth

import (
	"context"
	"errors"
	"log/slog"
	"ticket-booking/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(ctx context.Context, email string, password string) (string, error)
}

func NewService(jwt *JWTManager, userService user.Service, logger *slog.Logger) Service {
	return &authService{
		jwt:         jwt,
		userServise: userService,
		logger:      logger,
	}
}

type authService struct {
	jwt         *JWTManager
	userServise user.Service
	logger      *slog.Logger
}

func (s *authService) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := s.userServise.GetUserAuthByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	s.logger.Debug("passwords", "hashed", user.PasswordHash, "getted", password)
	ok := s.verifyPassword(user.PasswordHash, password)
	if !ok {
		return "", errors.New("Wrong password")
	}

	return s.generateToken(user.ID, string(user.Role))
}

func (s *authService) generateToken(userID int64, role string) (string, error) {
	return s.jwt.Generate(userID, role)
}

func (s *authService) verifyPassword(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
