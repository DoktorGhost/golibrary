package config

import (
	"sync"

	"github.com/caarlos0/env/v6"
)

var (
	once   sync.Once
	Config config
)

type config struct {
	LibraryPostgres DBConfig

	Secrets SecretConfig
}

type DBConfig struct {
	DbHost  string `env:"DB_HOST"`
	DbPort  string `env:"DB_PORT"`
	DbName  string `env:"DB_NAME"`
	DbLogin string `env:"DB_LOGIN"`
	DbPass  string `env:"DB_PASS"`
}

type SecretConfig struct {
	JWTSecret string `env:"SECRET_KEY_JWT"`
}

func LoadConfig() config {
	once.Do(func() {
		//считываем все переменны окружения в cfg
		if err := env.Parse(&Config); err != nil {
			panic(err)
		}
	})
	return Config

}
