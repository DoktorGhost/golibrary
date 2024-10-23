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
	LibraryPostgres DBConfig     `mapstructure:"LibraryPostgres"`
	Secrets         SecretConfig `mapstructure:"Secrets"`
	GrpcConfig      GrpcConfig   `mapstructure:"GrpcConfig"`
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

type GrpcConfig struct {
	UserHost string `mapstructure:"USER_HOST"`
	UserPort string `mapstructure:"USER_PORT"`
}

func LoadConfig() config {
	once.Do(func() {
		// Декодируем значения в структуру Config
		viper.BindEnv("LibraryPostgres.DB_HOST", "DB_HOST")
		viper.BindEnv("LibraryPostgres.DB_PORT", "DB_PORT")
		viper.BindEnv("LibraryPostgres.DB_NAME", "DB_NAME")
		viper.BindEnv("LibraryPostgres.DB_LOGIN", "DB_LOGIN")
		viper.BindEnv("LibraryPostgres.DB_PASS", "DB_PASS")

		viper.BindEnv("Secrets.SECRET_KEY_JWT", "SECRET_KEY_JWT")

		viper.BindEnv("GrpcConfig.USER_HOST", "USER_HOST")
		viper.BindEnv("GrpcConfig.USER_PORT", "USER_PORT")

		if err := viper.Unmarshal(&Config); err != nil {
			panic(fmt.Errorf("ошибка декодирования конфигурации: %w", err))
		}

	})

	return Config
}
