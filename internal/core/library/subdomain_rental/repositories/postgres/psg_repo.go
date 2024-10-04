package postgres

import (
	"database/sql"
)

type RentalRepository struct {
	db *sql.DB
}

func NewRentalPostgresRepository(db *sql.DB) *RentalRepository {
	return &RentalRepository{db: db}
}
