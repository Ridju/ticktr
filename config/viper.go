package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBSource            string        `mapstructure:"DB_SOURCE"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	AccessTokenKey      string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	DBDriver            string        `mapstructure:"DB_DRIVER"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
