package containers

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
	"path/filepath"
	"testing"
)

func SetupPostgresContainer(t *testing.T) (*pgxpool.Pool, func()) {
	ctx := context.Background()

	// Настройка контейнера PostgreSQL
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "test_user",
			"POSTGRES_PASSWORD": "test_pas",
			"POSTGRES_DB":       "test_db",
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections").WithOccurrence(2),
	}

	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatal(err)
	}

	// Получаем хост и порт контейнера
	host, err := postgresContainer.Host(ctx)
	if err != nil {
		t.Fatal(err)
	}

	port, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatal(err)
	}

	// Создаем строку подключения
	dsn := fmt.Sprintf("postgres://test_user:test_pas@%s:%s/test_db?sslmode=disable", host, port.Port())

	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		t.Fatal(err)
	}

	//migrationsDir, err := filepath.Abs("../../migrations")
	migrationsDir, err := filepath.Abs("../migrations")
	if err != nil {
		t.Fatal(err)
	}

	// Проверка существования директории
	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		t.Fatal(err)
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		t.Fatal(err)
	}

	if err := goose.Up(db, migrationsDir); err != nil {
		t.Fatal(err)
	}

	// Возвращаем базу данных и функцию очистки ресурсов
	return dbpool, func() {
		db.Close()
		postgresContainer.Terminate(ctx)
	}
}
