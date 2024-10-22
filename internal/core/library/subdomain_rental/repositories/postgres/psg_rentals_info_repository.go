package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/repositories/postgres/dao"
)

func (s *RentalRepository) CreateRentalsInfo(userID, bookID int) (int, error) {
	var id int
	query := `INSERT INTO rentals_info (user_id, book_id) VALUES ($1, $2) RETURNING id`
	err := s.db.QueryRow(context.Background(), query, userID, bookID).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("ошибка добавления записи: %v", err)
	}

	return id, nil
}

func (s *RentalRepository) GetRentalsInfoByID(id int) (dao.RentalsTable, error) {
	var result dao.RentalsTable
	var returnDate sql.NullTime
	query := `SELECT * FROM rentals_info WHERE id = $1`
	err := s.db.QueryRow(context.Background(), query, id).Scan(&result.ID, &result.UserID, &result.BookID, &result.RentalDate, &returnDate)

	if err != nil {
		if err == sql.ErrNoRows {
			return dao.RentalsTable{}, fmt.Errorf("запись с ID %d не найдена", id)
		}
		return dao.RentalsTable{}, fmt.Errorf("ошибка получения записи аренды: %v", err)
	}

	if returnDate.Valid {
		result.ReturnDate = returnDate.Time
	} else {
		result.ReturnDate = time.Time{}
	}

	return result, nil
}

func (s *RentalRepository) UpdateRentalsInfo(rentals dao.RentalsTable) error {
	query := `UPDATE rentals_info SET user_id = $1, book_id=$2, rental_date=$3, return_date=$4 WHERE id = $5`
	result, err := s.db.Exec(context.Background(), query, rentals.UserID, rentals.BookID, rentals.RentalDate, rentals.ReturnDate, rentals.ID)

	if err != nil {
		return fmt.Errorf("ошибка обновления записи аренды: %v", err)
	}

	// Проверяем, была ли обновлена хотя бы одна запись
	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("запись с ID %d не найдена", rentals.ID)
	}

	return nil
}

func (s *RentalRepository) DeleteRentalsInfo(id int) error {
	query := `DELETE FROM rentals_info WHERE id=$1`
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

func (s *RentalRepository) GetActiveRentals() (map[int][]int, error) {
	query := `SELECT user_id, book_id FROM rentals_info WHERE return_date IS NULL;`
	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	defer rows.Close()

	rentals := make(map[int][]int)

	for rows.Next() {
		var userID, bookID int

		err := rows.Scan(&userID, &bookID)
		if err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки: %v", err)
		}

		rentals[userID] = append(rentals[userID], bookID)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при чтении строк: %v", err)
	}

	return rentals, nil
}

func (s *RentalRepository) GetTopAuthors(days, limit int) ([]dao.TopAuthor, error) {
	period := fmt.Sprintf("%d days", days)

	// Формируем запрос
	query := fmt.Sprintf(`
			SELECT authors.name, COUNT(rentals_info.id) AS rental_count
			FROM authors
			JOIN books ON authors.id = books.author_id
			JOIN rentals_info ON books.id = rentals_info.book_id
			WHERE rentals_info.rental_date >= NOW() - INTERVAL '%s'
			GROUP BY authors.id
			ORDER BY rental_count DESC
			LIMIT $1;`, period)

	// Выполняем запрос, передавая только лимит
	rows, err := s.db.Query(context.Background(), query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []dao.TopAuthor
	for rows.Next() {
		var author dao.TopAuthor
		if err := rows.Scan(&author.Name, &author.CountRent); err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	return authors, nil
}
