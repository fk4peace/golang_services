package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	HttpServer
	Postgres
	Service
}

type Postgres struct {
	Username string `env:"POSTGRES_USERNAME"`
	Password string `env:"POSTGRES_PASSWORD"`
	Host     string `env:"POSTGRES_HOST"`
	Port     int    `env:"POSTGRES_PORT"`
	Database string `env:"POSTGRES_DATABASE"`
}

type HttpServer struct {
	Port int `env:"AUTH_SERVICE_HTTP_PORT"`
}

type Service struct {
	PasswordSalt     string `env:"AUTH_SERVICE_PASSWORD_SALT"`
	JwtAccessSecret  string `env:"AUTH_SERVICE_JWT_ACCESS_SECRET"`
	JwtRefreshSecret string `env:"AUTH_SERVICE_JWT_REFRESH_SECRET"`
}

func MustLoad() *Config {
	godotenv.Load("../.env")

	cfg := &Config{}

	if err := env.Parse(&cfg.HttpServer); err != nil {
		panic(err)
	}

	if err := env.Parse(&cfg.Postgres); err != nil {
		panic(err)
	}

	if err := env.Parse(&cfg.Service); err != nil {
		panic(err)
	}

	return cfg
}
