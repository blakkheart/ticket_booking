package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte("super-duper-secret-key")

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
	UserID int64  `json:"uid"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type JWTManager struct {
	secret []byte
	issuer string
	ttl    time.Duration
}

func NewJWTManager(secret string, issuer string, ttl time.Duration) *JWTManager {
	return &JWTManager{
		secret: []byte(secret),
		issuer: issuer,
		ttl:    ttl,
	}
}

func (j *JWTManager) Generate(userID int64, role string) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
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
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}
