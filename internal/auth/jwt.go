package auth

import (
	"errors"
	"net/http"
	"strings"
	"ticket-booking/internal/auth/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GetTokenFromPayload(r *http.Request) (string, error) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		return "", errors.New("error in payload")
	}

	payload := strings.TrimSpace(splitToken[1])

	return payload, nil

}

type Claims struct {
	UserID uuid.UUID `json:"uid"`
	Role   string    `json:"role"`
	jwt.RegisteredClaims
}

type JWTManager struct {
	secret     []byte
	issuer     string
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewJWTManager(
	secret string,
	issuer string,
	accessTTL time.Duration,
	refreshTTL time.Duration,
) *JWTManager {
	return &JWTManager{
		secret:     []byte(secret),
		issuer:     issuer,
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
	}
}

func (j *JWTManager) generateTokenWithTTL(
	userID uuid.UUID,
	role string,
	ttl time.Duration,
) (string, error) {
	expiresAt := time.Now().Add(ttl)
	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}

func (j *JWTManager) GenerateAccessToken(
	userID uuid.UUID,
	role string,
) (string, error) {
	accessToken, err := j.generateTokenWithTTL(userID, role, j.accessTTL)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (j *JWTManager) GenerateRefreshToken(
	userID uuid.UUID,
	role string,
) (string, error) {
	refreshToken, err := j.generateTokenWithTTL(userID, role, j.refreshTTL)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

func (j *JWTManager) GenerateTokenPair(
	userID uuid.UUID,
	role string,
) (*models.JWTTokens, error) {
	accessToken, err := j.GenerateAccessToken(userID, role)
	if err != nil {
		return nil, err
	}
	refreshToken, err := j.GenerateRefreshToken(userID, role)
	if err != nil {
		return nil, err
	}
	return &models.JWTTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (j *JWTManager) Parse(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&Claims{},
		func(t *jwt.Token) (any, error) {
			return j.secret, nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithExpirationRequired(),
		jwt.WithIssuer(j.issuer),
	)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}
