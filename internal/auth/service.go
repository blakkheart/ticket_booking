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
	GetTokenByUserID(ctx context.Context, userID int64) (*RefreshToken, error)
	UpdateTokenByUserID(ctx context.Context, newToken *RefreshToken) error
}

func NewService(repo Repository, jwt *JWTManager, userService user.Service, logger *slog.Logger) Service {
	return &authService{
		Repo:        repo,
		jwt:         jwt,
		userServise: userService,
		logger:      logger,
	}
}

type authService struct {
	Repo        Repository
	jwt         *JWTManager
	userServise user.Service
	logger      *slog.Logger
}

func (s *authService) Login(ctx context.Context, email string, password string) (*JWTTokens, error) {
	user, err := s.userServise.GetUserAuthByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
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

func (s *authService) GetTokenByUserID(ctx context.Context, userID int64) (*RefreshToken, error) {
	return s.Repo.GetTokenByUserID(ctx, userID)
}

func (s *authService) UpdateTokenByUserID(ctx context.Context, newToken *RefreshToken) error {
	return s.Repo.UpdateTokenByUserID(ctx, newToken)
}
