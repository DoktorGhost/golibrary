package repositories

import (
	"database/sql"
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/models"
)

func (s *PostgresRepository) CreateBook(book models.BookTable) (int, error) {
	var id int
	query := `INSERT INTO books (title, author_id) VALUES ($1, $2) RETURNING id`
	err := s.db.QueryRow(query, book.Title, book.AuthorID).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("ошибка добавления книги с названием '%s' и автором ID %d: %v", book.Title, book.AuthorID, err)
	}

	return id, nil
}

func (s *PostgresRepository) GetBookByID(id int) (models.BookTable, error) {
	var result models.BookTable
	query := `SELECT title, author_id FROM books WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&result.Title, &result.AuthorID)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.BookTable{}, fmt.Errorf("книга с ID %d не найдена", id)
		}
		return models.BookTable{}, fmt.Errorf("ошибка получения книги: %v", err)
	}
	result.ID = id

	return result, nil
}

func (s *PostgresRepository) UpdateBook(book models.BookTable) error {
	query := `UPDATE books SET title = $1, author_id = $2 WHERE id = $3`
	result, err := s.db.Exec(query, book.Title, book.AuthorID, book.ID)

	if err != nil {
		return fmt.Errorf("ошибка обновления книги: %v", err)
	}

	// Проверяем, была ли обновлена хотя бы одна запись
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения результата обновления: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("книга с ID %d не найдена", book.ID)
	}

	return nil
}

func (s *PostgresRepository) DeleteBook(id int) error {
	query := `DELETE FROM books WHERE id=$1`
	result, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("ошибка удаления книги: %v", err)
	}

	// Проверяем, была ли удалена хотя бы одна запись
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения результата удаления: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("книга с ID %d не найдена", id)
	}

	return nil
}

func (s *PostgresRepository) GetAllBooks() ([]models.BookTable, error) {
	query := `SELECT id, title, author_id FROM books;`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	defer rows.Close()

	var books []models.BookTable

	for rows.Next() {
		var book models.BookTable

		err := rows.Scan(&book.ID, &book.Title, &book.AuthorID)
		if err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки: %v", err)
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при чтении строк: %v", err)
	}

	return books, nil
}
