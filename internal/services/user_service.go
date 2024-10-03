package services

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/models"
	"github.com/DoktorGhost/golibrary/internal/services/crud"
)

type UserService struct {
	repo crud.UserRepository
}

func NewUserService(repo crud.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user models.UserTable) (int, error) {
	userID, err := s.repo.CreateUser(user)
	if err != nil {
		return 0, fmt.Errorf("ошибка создания пользователя: %v", err)
	}
	return userID, nil
}

func (s *UserService) GetUserById(id int) (models.UserTable, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return models.UserTable{}, err
	}
	return user, nil
}

func (s *UserService) GetUserByUsername(username string) (models.UserTable, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return models.UserTable{}, err
	}
	return user, nil
}

func (s *UserService) DeleteBook(id int) error {
	err := s.repo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateBook(user models.UserTable) error {
	err := s.repo.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}
