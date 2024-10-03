package repositories

import (
	"database/sql"
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/models"
	"time"
)

func (s *PostgresRepository) CreateRentalsInfo(userID, bookID int) (int, error) {
	var id int
	query := `INSERT INTO rentals_info (user_id, book_id) VALUES ($1, $2) RETURNING id`
	err := s.db.QueryRow(query, userID, bookID).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("ошибка добавления записи: %v", err)
	}

	return id, nil
}

func (s *PostgresRepository) GetRentalsInfoByID(id int) (models.RentalsTable, error) {
	var result models.RentalsTable
	var returnDate sql.NullTime
	query := `SELECT * FROM rentals_info WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&result.ID, &result.UserID, &result.BookID, &result.RentalDate, &returnDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.RentalsTable{}, fmt.Errorf("запись с ID %d не найдена", id)
		}
		return models.RentalsTable{}, fmt.Errorf("ошибка получения записи аренды: %v", err)
	}

	if returnDate.Valid {
		result.ReturnDate = returnDate.Time
	} else {
		result.ReturnDate = time.Time{} // Используем нулевое значение времени
	}

	return result, nil
}

func (s *PostgresRepository) UpdateRentalsInfo(rentals models.RentalsTable) error {
	query := `UPDATE rentals_info SET user_id = $1, book_id=$2, rental_date=$3, return_date=$4 WHERE id = $5`
	result, err := s.db.Exec(query, rentals.UserID, rentals.BookID, rentals.RentalDate, rentals.ReturnDate, rentals.ID)

	if err != nil {
		return fmt.Errorf("ошибка обновления записи аренды: %v", err)
	}

	// Проверяем, была ли обновлена хотя бы одна запись
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения результата обновления: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("запись с ID %d не найдена", rentals.ID)
	}

	return nil
}

func (s *PostgresRepository) DeleteRentalsInfo(id int) error {
	query := `DELETE FROM rentals_info WHERE id=$1`
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
