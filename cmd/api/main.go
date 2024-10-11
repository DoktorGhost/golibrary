package main

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/app"
	"github.com/DoktorGhost/golibrary/internal/delivery/controllers/handlers"
	"github.com/DoktorGhost/golibrary/internal/delivery/http/server"
	"github.com/DoktorGhost/golibrary/internal/enum"
	"github.com/DoktorGhost/golibrary/internal/metrics"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

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
	//инициализируем логгер
	logger, err := zaplogger.NewZapLogger()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logger.Sync()

	err = godotenv.Load("../../.env")
	if err != nil {
		logger.Error("ошибка загрузки файла .env", "error", err)
		return
	} else {
		logger.Info(".env файл успешно загружен")
	}

	pgsqlConnector, err := psg.InitStorage(config.LoadConfig().LibraryPostgres)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer pgsqlConnector.Close()
	logger.Info("соединение с БД установлено")

	cont := app.Init(pgsqlConnector)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx = context.WithValue(ctx, enum.UseCaseKeyProvider, cont.UseCaseProvider)

	r := handlers.SetupRoutes(cont.UseCaseProvider, logger)

	httpServer := server.NewHttpServer(r, ":8080")
	httpServer.Serve()

	err = cont.UseCaseProvider.DataUseCase.AddLibrary()
	if err != nil {
		logger.Error("Ошибка создания библиотеки:", "err", err)
	}

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	metrics.Init()

	//// Запуск HTTP-сервера
	//logger.Info("Запуск сервера на порту :8080")
	//if err := http.ListenAndServe(":8080", r); err != nil {
	//	logger.Fatal(err.Error())
	//}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)

	select {
	case killSignal := <-interrupt:
		logger.Info("Выключение сервера", "signal", killSignal)
	case err = <-httpServer.Notify():
		logger.Error("Ошибка сервера", "error", err)
	}

	httpServer.Shutdown()
}
