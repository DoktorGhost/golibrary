package main

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/app"
	"github.com/DoktorGhost/golibrary/internal/delivery/controllers/handlers"
	"github.com/DoktorGhost/golibrary/internal/delivery/http/server"
	"github.com/DoktorGhost/golibrary/internal/enum"
	"github.com/DoktorGhost/golibrary/internal/metrics"
	"github.com/DoktorGhost/golibrary/pkg/logger"
	"github.com/spf13/viper"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"context"
	"github.com/DoktorGhost/golibrary/config"
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
	log, err := logger.GetLogger()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer log.Sync()

	// Установка файла конфигурации .env
	viper.SetConfigFile("../../.env")

	// Чтение файла .env
	err = viper.ReadInConfig()
	if err != nil {
		log.Error("Ошибка загрузки файла .env", "error", err)
		return
	} else {
		log.Info(".env файл успешно загружен")
	}
	viper.AutomaticEnv()

	pgsqlConnector, err := psg.InitStorage(config.LoadConfig().LibraryPostgres)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer pgsqlConnector.Close()
	log.Info("соединение с БД установлено")

	cont := app.Init(pgsqlConnector)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx = context.WithValue(ctx, enum.UseCaseKeyProvider, cont.UseCaseProvider)

	r := handlers.SetupRoutes(cont.UseCaseProvider)

	httpServer := server.NewHttpServer(r, ":8080")
	httpServer.Serve()

	err = cont.UseCaseProvider.DataUseCase.AddLibrary()
	if err != nil {
		log.Error("Ошибка создания библиотеки:", "err", err)
	}

	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	metrics.Init()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)

	select {
	case killSignal := <-interrupt:
		log.Info("Выключение сервера", "signal", killSignal)
	case err = <-httpServer.Notify():
		log.Error("Ошибка сервера", "error", err)
	}

	httpServer.Shutdown()
}
