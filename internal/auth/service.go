package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log/slog"
	"ticket-booking/internal/auth/models"
	"ticket-booking/internal/repository"
	"ticket-booking/internal/user"

	"golang.org/x/crypto/bcrypt"
)

func hash(token string) string {
	h := sha256.Sum256([]byte(token))
	return hex.EncodeToString(h[:])
}

type Service interface {
	Login(ctx context.Context, email string, password string) (*models.JWTTokens, error)
	ParseToken(token string) (*Claims, error)
	GetTokenByUserID(ctx context.Context, userID int64) (*models.RefreshToken, error)
	UpdateTokenByUserID(ctx context.Context, newToken *models.RefreshToken) (*models.RefreshToken, error)
	GenerateTokenPair(userID int64, role string) (*models.JWTTokens, error)
	GenerateAccessToken(userID int64, role string) (string, error)
	GenerateRefreshToken(userID int64, role string) (string, error)
	GetTokenByHash(ctx context.Context, token string) (*models.RefreshToken, error)
}

func NewService(repo repository.AuthRepository, jwt *JWTManager, userService user.Service, logger *slog.Logger) Service {
	return &authService{
		Repo:        repo,
		jwt:         jwt,
		userServise: userService,
		logger:      logger,
	}
}

type authService struct {
	Repo        repository.AuthRepository
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

func (s *authService) Login(ctx context.Context, email string, password string) (*models.JWTTokens, error) {
	user, err := s.userServise.GetUserAuthByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	ok := s.verifyPassword(user.PasswordHash, password)
	if !ok {
		return nil, errors.New("invalid credentials")
	}

	tokens, err := s.GenerateTokenPair(user.ID, string(user.Role))

	claimsRefresh, err := s.ParseToken(tokens.RefreshToken)
	if err != nil {
		return nil, err
	}

	refreshHashed := hash(tokens.RefreshToken)

	_, errCreated := s.createOrUpdate(
		ctx,
		&models.RefreshToken{
			UserID:    user.ID,
			ExpiresAt: claimsRefresh.ExpiresAt.Time,
			Revoked:   false,
			TokenHash: refreshHashed,
		},
	)
	if errCreated != nil {
		return nil, errCreated
	}

	return tokens, nil

}

var RefreshTokenDuplicationError = errors.New("Duplication error")

func (s *authService) createOrUpdate(ctx context.Context, refreshToken *models.RefreshToken) (*models.RefreshToken, error) {

	_, errGet := s.Repo.GetTokenByUserID(ctx, refreshToken.UserID)

	if errGet != nil {
		createdToken, errCreate := s.Repo.Create(ctx, refreshToken)
		return createdToken, errCreate
	}

	return s.Repo.UpdateTokenByUserID(ctx, refreshToken)
}

func (s *authService) GenerateTokenPair(userID int64, role string) (*models.JWTTokens, error) {
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

func (s *authService) GetTokenByUserID(ctx context.Context, userID int64) (*models.RefreshToken, error) {
	return s.Repo.GetTokenByUserID(ctx, userID)
}

func (s *authService) GetTokenByHash(ctx context.Context, tokenHash string) (*models.RefreshToken, error) {
	return s.Repo.GetTokenByHash(ctx, tokenHash)
}

func (s *authService) UpdateTokenByUserID(ctx context.Context, newToken *models.RefreshToken) (*models.RefreshToken, error) {
	return s.Repo.UpdateTokenByUserID(ctx, newToken)
}
