package config

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/spf13/viper"
)

var allowedEnv = map[string]bool{
	"local":   true,
	"staging": true,
	"prod":    true,
}

type Config struct {
	App     AppConfig     `mapstructure:"app"`
	DB      DBConfig      `mapstructure:"db"`
	Server  ServerConfig  `mapstructure:"server"`
	Auth    AuthConfig    `mapstructure:"auth"`
	Payment PaymentConfig `mapstructure:"payment"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Env  string `mapstracture:"env"`
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
		"app.env",

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
		"auth.issuer",

		"payment.provider",
		"payment.secret",
	}

	for _, key := range keys {
		_ = v.BindEnv(key)
	}
}
func validateConfig() {
	if _, ok := allowedEnv[AppConfigs.App.Env]; !ok {
		keys := slices.Collect(maps.Keys(allowedEnv))
		result := strings.Join(keys, ", ")
		panic(fmt.Sprintf("ENV is not allowed. Choose from %s", result))
	}
	if AppConfigs.DB.Host == "" {
		panic("DB_HOST is not set")
	}

}
