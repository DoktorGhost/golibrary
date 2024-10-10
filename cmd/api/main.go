package main

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/app"
	"github.com/DoktorGhost/golibrary/internal/delivery/controllers/handlers"
	"github.com/DoktorGhost/golibrary/internal/enum"
	"log"
	"net/http"
	_ "net/http/pprof"

	"context"
	"github.com/DoktorGhost/golibrary/config"
	"github.com/DoktorGhost/golibrary/pkg/logger/zaplogger"
	"github.com/DoktorGhost/golibrary/pkg/storage/psg"
)

// @title Library
// @version 0.1.0
// @description Библиотека

// @securityDefinitions.apikey BearerAuth
// @type apiKey
// @name Authorization
// @in header
func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	//инициализируем логгер
	logger, err := zaplogger.NewZapLogger()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logger.Sync()

	conf, err := config.LoadConfig("../../.env", logger)
	//conf, err := config.LoadConfig(".env", logger)
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

	cont := app.Init(pgsqlConnector, conf)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx = context.WithValue(ctx, enum.UseCaseKeyProvider, cont.UseCaseProvider)

	r := handlers.SetupRoutes(cont.UseCaseProvider, logger)

	err = cont.UseCaseProvider.DataUseCase.AddLibrary()
	if err != nil {
		logger.Error("Ошибка создания библиотеки:", "err", err)
	}

	// Запуск HTTP-сервера
	logger.Info("Запуск сервера на порту :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Fatal(err.Error())
	}

}
