package config

import (
	"fmt"
)

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func (s *ServerConfig) GetAddress() string {
	return fmt.Sprintf(
		"%s:%d",
		s.Host,
		s.Port,
	)
}
