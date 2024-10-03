package repositories

import (
	"database/sql"
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/models"
	"github.com/lib/pq"
)

func (s *PostgresRepository) CreateUser(user models.UserTable) (int, error) {
	var id int
	query := `INSERT INTO users (username, password_hash, full_name) VALUES ($1, $2, $3) RETURNING id`
	err := s.db.QueryRow(query, user.Username, user.PasswordHash, user.FullName).Scan(&id)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return 0, fmt.Errorf("пользователь с таким именем уже существует")
		}
		return 0, fmt.Errorf("ошибка добавления записи: %v", err)
	}

	return id, nil
}

func (s *PostgresRepository) GetUserByID(id int) (models.UserTable, error) {
	var result models.UserTable
	query := `SELECT username, password_hash, full_name FROM users WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&result.Username, &result.PasswordHash, &result.FullName)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.UserTable{}, fmt.Errorf("пользователь с ID %d не найден", id)
		}
		return models.UserTable{}, fmt.Errorf("ошибка получения пользователя: %v", err)
	}
	result.ID = id
	return result, nil
}

func (s *PostgresRepository) UpdateUser(user models.UserTable) error {
	query := `UPDATE users SET username = $1, password_hash = $2, full_name = $3 WHERE id = $4`
	result, err := s.db.Exec(query, user.Username, user.PasswordHash, user.FullName, user.ID)

	if err != nil {
		return fmt.Errorf("ошибка обновления пользователя: %v", err)
	}

	// Проверяем, была ли обновлена хотя бы одна запись
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения результата обновления: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("пользователь с ID %d не найден", user.ID)
	}

	return nil
}

func (s *PostgresRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id=$1`
	result, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("ошибка удаления пользователя: %v", err)
	}

	// Проверяем, была ли удалена хотя бы одна запись
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения результата удаления: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("пользователь с ID %d не найден", id)
	}

	return nil
}

func (s *PostgresRepository) GetUserByUsername(username string) (models.UserTable, error) {
	var result models.UserTable
	query := `SELECT id, username, password_hash, full_name FROM users WHERE username = $1`
	err := s.db.QueryRow(query, username).Scan(&result.ID, &result.Username, &result.PasswordHash, &result.FullName)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.UserTable{}, fmt.Errorf("пользователь с username %s не найден", username)
		}
		return models.UserTable{}, fmt.Errorf("ошибка получения пользователя: %v", err)
	}
	return result, nil
}
