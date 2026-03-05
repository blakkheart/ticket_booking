package auth

import (
	"context"
	"errors"
	"log/slog"
	"ticket-booking/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(ctx context.Context, email string, password string) (*JWTTokens, error)
	ParseToken(token string) (*Claims, error)
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

func (s *authService) Login(ctx context.Context, email string, password string) (*JWTTokens, error) {
	user, err := s.userServise.GetUserAuthByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	s.logger.Debug("passwords", "hashed", user.PasswordHash, "getted", password)
	ok := s.verifyPassword(user.PasswordHash, password)
	if !ok {
		return nil, errors.New("Wrong password")
	}

	return s.generateTokenPair(user.ID, string(user.Role))
}

func (s *authService) generateTokenPair(userID int64, role string) (*JWTTokens, error) {
	return s.jwt.GenerateTokenPair(userID, role)
}

func (s *authService) verifyPassword(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}

func (s *authService) ParseToken(token string) (*Claims, error) {
	return s.jwt.Parse(token)
}
