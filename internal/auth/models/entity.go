package models

import (
	"time"

	"github.com/google/uuid"
)

type JWTTokens struct {
	RefreshToken string
	AccessToken  string
}

type RefreshToken struct {
	id         uuid.UUID
	UserID     uuid.UUID
	TokenHash  string
	ExpiresAt  time.Time
	Revoked    bool
	ReplacedBy *uuid.UUID
	Token      string
	Role       string
}

type AccessToken struct {
	id        uuid.UUID
	UserID    uuid.UUID
	Token     string
	ExpiresAt time.Time
	Role      string
}
