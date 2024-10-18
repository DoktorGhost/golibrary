package services

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/metrics"
	"time"

	"github.com/DoktorGhost/golibrary/internal/core/user/repositories/postgres/dao"
)

// UserRepository определяет методы для работы с пользователями
//
//go:generate mockgen -source=$GOFILE -destination=./mock_user.go -package=${GOPACKAGE}
type UserRepository interface {
	CreateUser(user dao.UserTable) (int, error)
	GetUserByID(id int) (dao.UserTable, error)
	UpdateUser(user dao.UserTable) error
	DeleteUser(id int) error
	GetUserByUsername(username string) (dao.UserTable, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user dao.UserTable) (int, error) {
	start := time.Now()

	userID, err := s.repo.CreateUser(user)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("CreateUser", duration)

	if err != nil {
		return 0, fmt.Errorf("ошибка создания пользователя: %v", err)
	}
	return userID, nil
}

func (s *UserService) GetUserById(id int) (dao.UserTable, error) {
	start := time.Now()

	user, err := s.repo.GetUserByID(id)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("GetUserById", duration)

	if err != nil {
		return dao.UserTable{}, fmt.Errorf("ошибка получения пользователя: %v", err)
	}
	return user, nil
}

func (s *UserService) GetUserByUsername(username string) (dao.UserTable, error) {
	start := time.Now()

	user, err := s.repo.GetUserByUsername(username)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("GetUserByUsername", duration)

	if err != nil {
		return dao.UserTable{}, err
	}
	return user, nil
}

func (s *UserService) DeleteUser(id int) error {
	start := time.Now()

	err := s.repo.DeleteUser(id)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("DeleteUser", duration)

	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateUser(user dao.UserTable) error {
	start := time.Now()

	err := s.repo.UpdateUser(user)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("UpdateUser", duration)

	if err != nil {
		return err
	}
	return nil
}
