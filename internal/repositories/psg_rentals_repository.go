package repositories

import (
	"database/sql"
	"fmt"
)

func (s *PostgresRepository) CreateRentals(bookID, rentalID int) error {
	query := `INSERT INTO rentals (id, rentals_id) VALUES ($1, $2)`
	err := s.db.QueryRow(query, bookID, rentalID)

	if err != nil {
		return fmt.Errorf("ошибка добавления записи: %v", err)
	}

	return nil
}

func (s *PostgresRepository) GetRentalsByID(id int) (int, error) {
	var rentals_id int
	query := `SELECT rentals_id FROM rentals WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&rentals_id)

	if err != nil {
		if err == sql.ErrNoRows {
			return -1, fmt.Errorf("запись с ID %d не найдена", id)
		}
		return -2, fmt.Errorf("ошибка получения записи аренды: %v", err)
	}

	return rentals_id, nil
}

func (s *PostgresRepository) UpdateRentals(id, rentals_id int) error {
	query := `UPDATE rentals SET rentals_id=$1 WHERE id = $2`
	result, err := s.db.Exec(query, rentals_id, id)

	if err != nil {
		return fmt.Errorf("ошибка обновления записи аренды: %v", err)
	}

	// Проверяем, была ли обновлена хотя бы одна запись
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения результата обновления: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("запись с ID %d не найдена", id)
	}

	return nil
}

func (s *PostgresRepository) DeleteRentals(id int) error {
	query := `DELETE FROM rentals WHERE id=$1`
	result, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("ошибка удаления записи: %v", err)
	}

	// Проверяем, была ли удалена хотя бы одна запись
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения результата удаления: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("запись с ID %d не найдена", id)
	}

	return nil
}
