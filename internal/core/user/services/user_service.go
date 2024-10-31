package services

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/core/user/entities"
	grpcrepo "github.com/DoktorGhost/golibrary/internal/core/user/repositories/grpc"
	"github.com/DoktorGhost/golibrary/internal/metrics"
	"time"
)

// UserRepository определяет методы для работы с пользователями
//
//go:generate mockgen -source=$GOFILE -destination=./mock_user.go -package=${GOPACKAGE}
type UsersRepository interface {
	Register(entities.RegisterData) (string, error)
	Login(entities.Login) (entities.UserTable, error)
	GetUserById(int) (string, error)
}

type UserService struct {
	repo grpcrepo.UsersRepository
}

func NewUserService(repo grpcrepo.UsersRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(userData entities.RegisterData) (string, error) {
	start := time.Now()

	userID, err := s.repo.Register(userData)

	duration := time.Since(start).Seconds()
	metrics.TrackExternalAPIDuration("ClientsService", "Register", duration)

	if err != nil {
		return "", fmt.Errorf("ошибка создания пользователя: %v", err)
	}
	return userID, nil
}

func (s *UserService) Login(userData entities.Login) (entities.UserTable, error) {
	start := time.Now()

	user, err := s.repo.Login(userData)

	duration := time.Since(start).Seconds()
	metrics.TrackExternalAPIDuration("ClientsService", "Login", duration)

	if err != nil {
		return entities.UserTable{}, fmt.Errorf("ошибка получения пользователя: %v", err)
	}
	return user, nil
}

func (s *UserService) GetUserById(userID int) (string, error) {
	start := time.Now()

	username, err := s.repo.GetUserById(userID)

	duration := time.Since(start).Seconds()
	metrics.TrackExternalAPIDuration("ClientsService", "GetUserById", duration)

	if err != nil {
		return "", fmt.Errorf("ошибка получения пользователя: %v", err)
	}

	return username, nil
}
