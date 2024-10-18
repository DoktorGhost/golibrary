package psg

import (
	"context"
	"fmt"
	"github.com/DoktorGhost/golibrary/pkg/users/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitStorage(conf *config.Config, schema []byte) (*pgxpool.Pool, error) {
	login := conf.DB_login
	password := conf.DB_pass
	host := conf.DB_host
	port := conf.DB_port
	dbname := conf.DB_name

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", login, password, host, port, dbname)

	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к бд: %v", err)
	}

	err = dbpool.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса к бд: %v", err)
	}

	// Выполняем SQL запросы из файла
	_, err = dbpool.Exec(context.Background(), string(schema))
	if err != nil {
		return nil, fmt.Errorf("ошибка применения схемы: %v", err)
	}

	return dbpool, nil
}
