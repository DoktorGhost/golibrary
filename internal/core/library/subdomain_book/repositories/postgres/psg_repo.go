package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookRepository struct {
	db *pgxpool.Pool
}

func NewBookPostgresRepository(db *pgxpool.Pool) *BookRepository {
	return &BookRepository{db: db}
}
