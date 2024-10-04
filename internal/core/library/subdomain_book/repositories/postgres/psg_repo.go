package postgres

import (
	"database/sql"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookPostgresRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}
