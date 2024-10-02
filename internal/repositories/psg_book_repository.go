package repositories

import (
	"database/sql"
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/models"
)

func (s *PostgresRepository) CreateBook(book models.Book) (int, error) {
	var id int
	_, err := s.GetAuthorByID(book.AuthorID)
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO books (title, author_id) VALUES ($1, $2) RETURNING id`
	err = s.db.QueryRow(query, book.Title, book.AuthorID).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("ошибка добавления записи: %v", err)
	}

	return id, nil
}

func (s *PostgresRepository) GetBookByID(id int) (models.Book, error) {
	var result models.Book
	query := `SELECT title, author_id FROM books WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&result.Title, &result.AuthorID)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Book{}, fmt.Errorf("книга с ID %d не найдена", id)
		}
		return models.Book{}, fmt.Errorf("ошибка получения книги: %v", err)
	}
	result.ID = id
	return result, nil
}

func (s *PostgresRepository) UpdateBook(book models.Book) error {
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
