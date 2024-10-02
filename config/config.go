package config

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/logger"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST  string `env:"DB_HOST"`
	DB_PORT  string `env:"DB_PORT"`
	DB_NAME  string `env:"DB_NAME"`
	DB_LOGIN string `env:"DB_LOGIN"`
	DB_PASS  string `env:"DB_PASS"`
}

func LoadConfig(path string, logger logger.Logger) (*Config, error) {

	if path != "" {
		err := godotenv.Load(path)
		if err != nil {
			logger.Error("Ошибка загрузки файла .env", "error", err)
			return nil, fmt.Errorf("Ошибка загрузки файла .env: %v", err)
		} else {
			logger.Info(".env файл успешно загружен")
		}
	}

	config := &Config{}
	//считываем все переменны окружения в cfg
	if err := env.Parse(config); err != nil {
		logger.Error("Ошибка загрузки конфигурации", "error", err)
		return nil, fmt.Errorf("Ошибка загрузки конфигурации: %v", err)
	}

	logger.Info("Конфигурация успешно загружена")
	return config, nil
}
