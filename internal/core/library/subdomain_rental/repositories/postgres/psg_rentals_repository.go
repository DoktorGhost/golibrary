package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
)

func (s *RentalRepository) CreateRentals(bookID int) error {
	query := `INSERT INTO library.rentals (id) VALUES ($1)`
	_, err := s.db.Exec(context.Background(), query, bookID)
	if err != nil {
		return fmt.Errorf("ошибка добавления записи: %v", err)
	}
	return nil
}

func (s *RentalRepository) GetRentalsByID(id int) (int, error) {
	var rentals_id *int
	query := `SELECT rentals_id FROM library.rentals WHERE id = $1`
	err := s.db.QueryRow(context.Background(), query, id).Scan(&rentals_id)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("запись с ID %d не найдена", id)
		}
		return 0, fmt.Errorf("ошибка получения записи аренды: %v", err)
	}
	if rentals_id == nil {
		return 0, nil
	}

	return *rentals_id, nil
}

func (s *RentalRepository) UpdateRentals(id, rentals_id int) error {
	var query string
	var result pgconn.CommandTag
	var err error

	// Если rentals_id равен 0, используем NULL в запросе
	if rentals_id == 0 {
		query = `UPDATE library.rentals SET rentals_id = NULL WHERE id = $1`
		result, err = s.db.Exec(context.Background(), query, id)
	} else {
		query = `UPDATE library.rentals SET rentals_id = $1 WHERE id = $2`
		result, err = s.db.Exec(context.Background(), query, rentals_id, id)
	}

	if err != nil {
		return fmt.Errorf("ошибка обновления записи аренды: %v", err)
	}

	// Проверяем, была ли обновлена хотя бы одна запись
	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("запись с ID %d не найдена", id)
	}

	return nil
}

func (s *RentalRepository) DeleteRentals(id int) error {
	query := `DELETE FROM library.rentals WHERE id=$1`
	result, err := s.db.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("ошибка удаления записи: %v", err)
	}

	// Проверяем, была ли удалена хотя бы одна запись
	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("запись с ID %d не найдена", id)
	}

	return nil
}
