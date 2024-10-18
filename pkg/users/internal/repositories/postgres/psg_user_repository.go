package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DoktorGhost/golibrary/users/internal/repositories/postgres/dao"
	"github.com/lib/pq"
)

func (s *UsersRepository) CreateUser(user dao.UserTable) (int, error) {
	var id int
	query := `INSERT INTO users (username, password_hash, full_name) VALUES ($1, $2, $3) RETURNING id`
	err := s.db.QueryRow(context.Background(), query, user.Username, user.PasswordHash, user.FullName).Scan(&id)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return 0, fmt.Errorf("пользователь с таким именем уже существует")
		}
		return 0, fmt.Errorf("ошибка добавления записи: %v", err)
	}

	return id, nil
}

func (s *UsersRepository) GetUserByID(id int) (dao.UserTable, error) {
	var result dao.UserTable
	query := `SELECT username, password_hash, full_name FROM users WHERE id = $1`
	err := s.db.QueryRow(context.Background(), query, id).Scan(&result.Username, &result.PasswordHash, &result.FullName)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dao.UserTable{}, fmt.Errorf("пользователь с ID %d не найден", id)
		}
		return dao.UserTable{}, fmt.Errorf("[PostgreSQL] Error: %v", err)
	}
	result.ID = id
	return result, nil
}

func (s *UsersRepository) UpdateUser(user dao.UserTable) error {
	query := `UPDATE users SET username = $1, password_hash = $2, full_name = $3 WHERE id = $4`
	result, err := s.db.Exec(context.Background(), query, user.Username, user.PasswordHash, user.FullName, user.ID)

	if err != nil {
		return fmt.Errorf("ошибка обновления пользователя: %v", err)
	}

	// Проверяем, была ли обновлена хотя бы одна запись
	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("пользователь с ID %d не найден", user.ID)
	}

	return nil
}

func (s *UsersRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id=$1`
	result, err := s.db.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("ошибка удаления пользователя: %v", err)
	}

	// Проверяем, была ли удалена хотя бы одна запись
	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("пользователь с ID %d не найден", id)
	}

	return nil
}

func (s *UsersRepository) GetUserByUsername(username string) (dao.UserTable, error) {
	var result dao.UserTable
	query := `SELECT id, username, password_hash, full_name FROM users WHERE username = $1`
	err := s.db.QueryRow(context.Background(), query, username).Scan(&result.ID, &result.Username, &result.PasswordHash, &result.FullName)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dao.UserTable{}, fmt.Errorf("пользователь с username %s не найден", username)
		}
		return dao.UserTable{}, fmt.Errorf("ошибка получения пользователя: %v", err)
	}
	return result, nil
}
