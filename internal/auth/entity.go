package auth

import "time"

type JWTTokens struct {
	RefreshToken string
	AccessToken  string
}

type RefreshToken struct {
	UserID     int64
	TokenHash  string
	ExpiresAt  time.Time
	Revoked    bool
	ReplacedBy *int64
	Token      string
	Role       string
}

type AccessToken struct {
	UserID    int64
	Token     string
	ExpiresAt time.Time
	Role      string
}
