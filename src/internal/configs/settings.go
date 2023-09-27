package configs

import (
	"errors"
	"github.com/spf13/viper"
)

type PgConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	DbName   string
	SslMode  string
}

type Config struct {
	Postgres PgConfig
}

func NewConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigType("yml")
	v.SetConfigName("config")
	v.AddConfigPath("internal/configs/")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		var configNotFound viper.ConfigFileNotFoundError
		switch {
		case errors.As(err, &configNotFound):
			return nil, errors.New("config.yml file not found")
		}
	}

	var config Config
	err = v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
