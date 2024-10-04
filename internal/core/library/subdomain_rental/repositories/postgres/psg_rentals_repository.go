package postgres

import (
	"database/sql"
	"fmt"
)

func (s *RentalRepository) CreateRentals(bookID int) error {
	query := `INSERT INTO library.rentals (id) VALUES ($1)`
	_, err := s.db.Exec(query, bookID)
	if err != nil {
		return fmt.Errorf("ошибка добавления записи: %v", err)
	}
	return nil
}

func (s *RentalRepository) GetRentalsByID(id int) (int, error) {
	var rentals_id int
	query := `SELECT rentals_id FROM library.rentals WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&rentals_id)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("запись с ID %d не найдена", id)
		}
		return 0, fmt.Errorf("ошибка получения записи аренды: %v", err)
	}

	return rentals_id, nil
}

func (s *RentalRepository) UpdateRentals(id, rentals_id int) error {
	var query string
	var result sql.Result
	var err error

	// Если rentals_id равен 0, используем NULL в запросе
	if rentals_id == 0 {
		query = `UPDATE library.rentals SET rentals_id = NULL WHERE id = $1`
		result, err = s.db.Exec(query, id)
	} else {
		query = `UPDATE library.rentals SET rentals_id = $1 WHERE id = $2`
		result, err = s.db.Exec(query, rentals_id, id)
	}

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

func (s *RentalRepository) DeleteRentals(id int) error {
	query := `DELETE FROM library.rentals WHERE id=$1`
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
