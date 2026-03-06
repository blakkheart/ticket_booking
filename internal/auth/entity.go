package auth

import "time"

type JWTTokens struct {
	RefreshToken string
	AccessToken  string
}

type RefreshToken struct {
	UserID    int64
	TokenHash string
	ExpiresAt time.Time
	Revoked   bool
}
