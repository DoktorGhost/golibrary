package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var (
	once   sync.Once
	Config config
)

type config struct {
	LibraryPostgres DBConfig
	Secrets         SecretConfig
}

type DBConfig struct {
	DbHost  string `mapstructure:"DB_HOST"`
	DbPort  string `mapstructure:"DB_PORT"`
	DbName  string `mapstructure:"DB_NAME"`
	DbLogin string `mapstructure:"DB_LOGIN"`
	DbPass  string `mapstructure:"DB_PASS"`
}

type SecretConfig struct {
	JWTSecret string `mapstructure:"SECRET_KEY_JWT"`
}

func LoadConfig() config {
	once.Do(func() {
		// Декодируем значения в структуру Config
		if err := viper.Unmarshal(&Config.LibraryPostgres); err != nil {
			panic(fmt.Errorf("ошибка декодирования конфигурации: %w", err))
		}
		if err := viper.Unmarshal(&Config.Secrets); err != nil {
			panic(fmt.Errorf("ошибка декодирования конфигурации: %w", err))
		}
	})

	return Config
}
