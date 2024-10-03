package repositories

import (
	"database/sql"
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/models"
)

func (s *PostgresRepository) CreateAuthor(name string) (int, error) {
	var id int
	query := `INSERT INTO authors (name) VALUES ($1) RETURNING id`
	err := s.db.QueryRow(query, name).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("ошибка добавления записи: %v", err)
	}

	return id, nil
}

func (s *PostgresRepository) GetAuthorByID(id int) (models.AuthorTable, error) {
	var result models.AuthorTable
	query := `SELECT name FROM authors WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&result.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.AuthorTable{}, fmt.Errorf("автор с ID %d не найден", id)
		}
		return models.AuthorTable{}, fmt.Errorf("ошибка получения автора: %v", err)
	}
	result.ID = id
	return result, nil
}

func (s *PostgresRepository) UpdateAuthor(author models.AuthorTable) error {
	query := `UPDATE authors SET name = $1 WHERE id = $2`
	result, err := s.db.Exec(query, author.Name, author.ID)

	if err != nil {
		return fmt.Errorf("ошибка обновления автора: %v", err)
	}

	// Проверяем, была ли обновлена хотя бы одна запись
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения результата обновления: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("автор с ID %d не найден", author.ID)
	}

	return nil
}

func (s *PostgresRepository) DeleteAuthor(id int) error {
	query := `DELETE FROM authors WHERE id=$1`
	result, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("ошибка удаления автора: %v", err)
	}

	// Проверяем, была ли удалена хотя бы одна запись
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения результата удаления: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("автор с ID %d не найден", id)
	}

	return nil
}

func (s *PostgresRepository) GetAllAuthors() ([]models.AuthorTable, error) {
	query := `SELECT id, name FROM authors;`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	defer rows.Close()

	var authors []models.AuthorTable

	for rows.Next() {
		var author models.AuthorTable

		err := rows.Scan(&author.ID, &author.Name)
		if err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки: %v", err)
		}

		authors = append(authors, author)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при чтении строк: %v", err)
	}

	return authors, nil
}
