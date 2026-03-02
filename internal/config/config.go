package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App    AppConfig    `mapstructure:"app"`
	DB     DBConfig     `mapstructure:"db"`
	Server ServerConfig `mapstructure:"server"`
	Auth   AuthConfig   `mapstructure:"auth"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
}

var AppConfigs Config

func InitConfigs() {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")

	_ = v.ReadInConfig() // dosen't work without yaml but we get env from enviroment itself

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	bindEnvs(v)

	if err := v.Unmarshal(&AppConfigs); err != nil {
		panic(fmt.Errorf("cannot unmarshal config: %w", err))
	}

	validateConfig()

}

func bindEnvs(v *viper.Viper) {
	keys := []string{
		"app.name",

		"db.host",
		"db.port",
		"db.user",
		"db.name",
		"db.password",
		"db.sslmode",

		"server.host",
		"server.port",

		"auth.secret_key",
		"auth.token_duration",
	}

	for _, key := range keys {
		_ = v.BindEnv(key)
	}
}
func validateConfig() {
	if AppConfigs.DB.Host == "" {
		panic("DB_HOST is not set")
	}
}
