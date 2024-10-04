package postgres

import (
	"database/sql"
)

type UsersRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db: db}
}
