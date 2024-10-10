package config

import (
	"fmt"

	"github.com/DoktorGhost/golibrary/pkg/logger"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	DbHost    string `env:"DB_HOST"`
	DbPort    string `env:"DB_PORT"`
	DbName    string `env:"DB_NAME"`
	DbLogin   string `env:"DB_LOGIN"`
	DbPass    string `env:"DB_PASS"`
	JWTSecret string `env:"SECRET_KEY_JWT"`
}

func LoadConfig(path string, logger logger.Logger) (*Config, error) {

	if path != "" {
		err := godotenv.Load(path)
		if err != nil {
			logger.Error("Ошибка загрузки файла .env", "error", err)
			return nil, fmt.Errorf("ошибка загрузки файла .env: %v", err)
		} else {
			logger.Info(".env файл успешно загружен")
		}
	}

	config := &Config{}
	//считываем все переменны окружения в cfg
	if err := env.Parse(config); err != nil {
		logger.Error("Ошибка загрузки конфигурации", "error", err)
		return nil, fmt.Errorf("ошибка загрузки конфигурации: %v", err)
	}

	logger.Info("Конфигурация успешно загружена")
	return config, nil
}
