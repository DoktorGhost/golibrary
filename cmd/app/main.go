package main

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/app"
	"github.com/DoktorGhost/golibrary/internal/delivery/controllers/handlers"
	"github.com/DoktorGhost/golibrary/internal/delivery/grpc/client"
	"github.com/DoktorGhost/golibrary/internal/delivery/http/server"
	"github.com/DoktorGhost/golibrary/internal/enum"
	"github.com/DoktorGhost/golibrary/internal/metrics"
	"github.com/DoktorGhost/platform/logger"
	"github.com/DoktorGhost/platform/storage/psg"
	"github.com/spf13/viper"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"context"
	"github.com/DoktorGhost/golibrary/config"
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
	viper.SetConfigFile(".env")

	// Чтение файла .env
	err = viper.ReadInConfig()
	if err != nil {
		log.Debug("Ошибка загрузки файла .env", "error", err)
	} else {
		log.Info(".env файл успешно загружен")
	}

	viper.AutomaticEnv()

	config.LoadConfig()

	//соединение с БД
	pgsqlConnector, err := psg.InitStorage(psg.DBConfig{
		config.LoadConfig().LibraryPostgres.DbHost,
		config.LoadConfig().LibraryPostgres.DbPort,
		config.LoadConfig().LibraryPostgres.DbName,
		config.LoadConfig().LibraryPostgres.DbLogin,
		config.LoadConfig().LibraryPostgres.DbPass,
	})

	if err != nil {
		log.Error(err.Error())
		return
	}
	defer pgsqlConnector.Close()
	log.Info("соединение с БД установлено")

	//grpc client
	userClient, conn := client.InitUserClient()
	defer conn.Close()

	cont := app.Init(pgsqlConnector, userClient)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx = context.WithValue(ctx, enum.UseCaseKeyProvider, cont.UseCaseProvider)

	r := handlers.SetupRoutes(cont.UseCaseProvider)

	//старт сервера
	httpServer := server.NewHttpServer(r, ":8080")
	httpServer.Serve()

	err = cont.UseCaseProvider.DataUseCase.AddLibrary()
	if err != nil {
		log.Error("Ошибка создания библиотеки:", "err", err)
	}

	//инициализация метрик
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
