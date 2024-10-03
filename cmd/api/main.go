package main

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/config"
	"github.com/DoktorGhost/golibrary/internal/logger/zaplogger"
	"github.com/DoktorGhost/golibrary/internal/repositories"
	"github.com/DoktorGhost/golibrary/internal/services"
	"github.com/DoktorGhost/golibrary/pkg/storage/psg"
)

func main() {
	//инициализируем логгер
	logger, err := zaplogger.NewZapLogger()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logger.Sync()

	conf, err := config.LoadConfig("../../.env", logger)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	pgsqlConnector, err := psg.InitStorage(conf)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Info("соединение с БД установлено")

	crudRepo := repositories.NewPostgresRepository(pgsqlConnector.DB)

	_ = services.NewAuthorService(crudRepo)

}
