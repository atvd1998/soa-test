package config

import (
	"time"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	App struct {
		HTTPAddr     string        `env:"HTTP_ADDR" envDefault:"localhost:8081"`
		StartTimeout time.Duration `env:"APP_START_TIMEOUT" envDefault:"1m"`
		StopTimeout  time.Duration `env:"APP_STOP_TIMEOUT" envDefault:"1m"`
	}
	Postgresql struct {
		Host            string        `env:"POSTGRESQL_HOST" envDefault:"3.1.28.125"`
		Port            string        `env:"POSTGRESQL_PORT" envDefault:"5432"`
		Username        string        `env:"POSTGRESQL_USERNAME" envDefault:"postgres"`
		Password        string        `env:"POSTGRESQL_PASSWORD" envDefault:"M1sIWvQ2D4MfWke7ReSt2IFHVPRXtpp6"`
		DbName          string        `env:"POSTGRESQL_DB_NAME" envDefault:"backend_test"`
		SSLMode         string        `env:"POSTGRESQL_SSL_MODE" envDefault:"disable"`
		MaxIdleConns    int           `env:"POSTGRESQL_MAX_IDLE_CONNS" envDefault:"10"`
		MaxOpenConns    int           `env:"POSTGRESQL_MAX_OPEN_CONNS" envDefault:"100"`
		ConnMaxLifetime time.Duration `env:"POSTGRESQL_CONN_MAX_LIFE_TIME" envDefault:"5m"`
	}
}

func Load() (*Config, error) {
	var conf Config
	if err := env.Parse(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

func MustLoad() *Config {
	conf, err := Load()
	if err != nil {
		panic(err)
	}
	return conf
}
