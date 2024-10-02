package main

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/config"
	"github.com/DoktorGhost/golibrary/internal/logger/zaplogger"
	"github.com/DoktorGhost/golibrary/pkg/storage/psg"
)

func main() {
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

	_, err = psg.InitStorage(conf)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Info("соединение с БД установлено")

}
