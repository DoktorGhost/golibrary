package repositories

import (
	"database/sql"
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/models"
	"github.com/lib/pq"
)

func (s *PostgresRepository) CreateUser(user models.User) (int, error) {
	var id int
	query := `INSERT INTO users (username, password_hash, fio) VALUES ($1, $2, $3) RETURNING id`
	err := s.db.QueryRow(query, user.Username, user.Password, user.FIO).Scan(&id)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return 0, fmt.Errorf("пользователь с таким именем уже существует")
		}
		return 0, fmt.Errorf("ошибка добавления записи: %v", err)
	}

	return id, nil
}

func (s *PostgresRepository) GetUserByID(id int) (models.User, error) {
	var result models.User
	query := `SELECT username, password_hash, fio FROM users WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&result.Username, &result.Password, &result.FIO)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("пользователь с ID %d не найден", id)
		}
		return models.User{}, fmt.Errorf("ошибка получения пользователя: %v", err)
	}
	result.ID = id
	return result, nil
}

func (s *PostgresRepository) UpdateUser(user models.User) error {
	query := `UPDATE users SET username = $1, password_hash = $2, fio = $3 WHERE id = $4`
	result, err := s.db.Exec(query, user.Username, user.Password, user.FIO, user.ID)

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
