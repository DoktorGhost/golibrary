package main

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/app"

	"github.com/DoktorGhost/golibrary/config"
	"github.com/DoktorGhost/golibrary/pkg/logger/zaplogger"
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

	app.Init(pgsqlConnector)

	/*
		repositoryProvider := providers.NewRepositoryProvider(pgsqlConnector.DB)
		repositoryProvider.RegisterDependencies()

		serviceProvider := providers.NewServiceProvider()
		serviceProvider.RegisterDependencies(repositoryProvider)

		useCaseProvider := providers.NewUseCaseProvider()
		useCaseProvider.RegisterDependencies(serviceProvider)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctx = context.WithValue(ctx, enum.UseCaseKeyProvider, useCaseProvider)

	*/

}
