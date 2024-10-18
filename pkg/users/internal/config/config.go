package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	DB_host       string `env:"DB_HOST"`
	DB_port       string `env:"DB_PORT"`
	DB_name       string `env:"DB_NAME"`
	DB_login      string `env:"DB_LOGIN"`
	DB_pass       string `env:"DB_PASS"`
	Provider_port string `env:"PROVIDER_PORT"`
}

func LoadConfig() (*Config, error) {
	config := &Config{}
	//считываем все переменны окружения в cfg
	if err := env.Parse(config); err != nil {
		return nil, fmt.Errorf("ошибка загрузки конфигурации: %v", err)
	}
	return config, nil
}
