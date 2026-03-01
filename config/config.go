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

	serverErr := viper.Unmarshal(&ServerConfig)
	if serverErr != nil {
		panic(fmt.Errorf("No correct Server structure was provided: %w", dbErr))
	}

	authErr := viper.Unmarshal(&AuthConfig)
	if authErr != nil {
		panic(fmt.Errorf("No correct Auth structure was provided: %w", dbErr))
	}

}
