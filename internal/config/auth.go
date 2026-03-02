package config

import "time"

type AuthConfig struct {
	SecretKey     string        `mapstructure:"secret_key"`
	TokenDuration time.Duration `mapstructure:"token_duration"`
}
