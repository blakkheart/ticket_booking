package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func ReadConfigs() {
	viper.SetConfigName("config")

	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	dbErr := viper.Unmarshal(&DBConfig)
	if dbErr != nil {
		panic(fmt.Errorf("No correct DB structure was provided: %w", dbErr))
	}
}
