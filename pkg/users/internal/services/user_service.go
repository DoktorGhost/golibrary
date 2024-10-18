package services

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/pkg/users/internal/repositories/postgres/dao"
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
	userID, err := s.repo.CreateUser(user)
	if err != nil {
		return 0, fmt.Errorf("ошибка создания пользователя: %v", err)
	}
	return userID, nil
}

func (s *UserService) GetUserById(id int) (dao.UserTable, error) {
	user, err := s.repo.GetUserByID(id)

	if err != nil {
		return dao.UserTable{}, fmt.Errorf("ошибка получения пользователя: %v", err)
	}
	return user, nil
}

func (s *UserService) GetUserByUsername(username string) (dao.UserTable, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return dao.UserTable{}, err
	}
	return user, nil
}

func (s *UserService) DeleteUser(id int) error {
	err := s.repo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateUser(user dao.UserTable) error {
	err := s.repo.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}
