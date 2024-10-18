package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type RentalRepository struct {
	db *pgxpool.Pool
}

func NewRentalPostgresRepository(db *pgxpool.Pool) *RentalRepository {
	return &RentalRepository{db: db}
}
