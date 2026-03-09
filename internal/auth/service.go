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
	GenerateTokenPair(userID int64, role string) (*JWTTokens, error)
	GenerateAccessToken(userID int64, role string) (string, error)
	GenerateRefreshToken(userID int64, role string) (string, error)
	GetTokenByHash(ctx context.Context, token string) (*RefreshToken, error)
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

func (s *authService) GenerateAccessToken(userID int64, role string) (string, error) {
	return s.jwt.GenerateAccessToken(userID, role)
}

func (s *authService) GenerateRefreshToken(userID int64, role string) (string, error) {
	return s.jwt.GenerateRefreshToken(userID, role)
}

func (s *authService) Login(ctx context.Context, email string, password string) (*JWTTokens, error) {
	user, err := s.userServise.GetUserAuthByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	ok := s.verifyPassword(user.PasswordHash, password)
	if !ok {
		return nil, errors.New("invalid credentials")
	}

	return s.GenerateTokenPair(user.ID, string(user.Role))
}

func (s *authService) GenerateTokenPair(userID int64, role string) (*JWTTokens, error) {
	return s.jwt.GenerateTokenPair(userID, role)
}

func (s *authService) verifyPassword(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}

func (s *authService) HashToken(token string) (string, error) {
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedToken), nil
}

func (s *authService) CompareTokenHashes(token string, hashedToken string) (bool, error) {
	incomingTokenHashed, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}
	return string(incomingTokenHashed) == hashedToken, nil
}

func (s *authService) ParseToken(token string) (*Claims, error) {
	return s.jwt.Parse(token)
}

func (s *authService) GetTokenByUserID(ctx context.Context, userID int64) (*RefreshToken, error) {
	return s.Repo.GetTokenByUserID(ctx, userID)
}

func (s *authService) GetTokenByHash(ctx context.Context, tokenHash string) (*RefreshToken, error) {
	return s.Repo.GetTokenByHash(ctx, tokenHash)
}

func (s *authService) UpdateTokenByUserID(ctx context.Context, newToken *RefreshToken) error {
	return s.Repo.UpdateTokenByUserID(ctx, newToken)
}
