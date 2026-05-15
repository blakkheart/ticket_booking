package auth

import "errors"

var (
	ErrTokenExpired              = errors.New("Token expired")
	RefreshTokenDuplicationError = errors.New("Duplication error")
)
