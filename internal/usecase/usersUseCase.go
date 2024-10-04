package usecase

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/models"
	"github.com/DoktorGhost/golibrary/internal/services"
	"github.com/DoktorGhost/golibrary/pkg/validator"
	"golang.org/x/crypto/bcrypt"
)

type UsersUseCase struct {
	userService services.UserService
}

func NewUsersUseCase(userService services.UserService) *UsersUseCase {
	return &UsersUseCase{userService: userService}
}

func (uc *UsersUseCase) AddUser(userData models.RegisterData) (int, error) {
	// Проверка, существует ли пользователь с таким именем
	_, err := uc.userService.GetUserByUsername(userData.Username)
	if err == nil {
		return -1, fmt.Errorf("пользователь с таким Username уже существует")
	}

	// Валидация данных пользователя
	fullName, err := validator.Valid(userData.Name, userData.Surname, userData.Patronymic)
	if err != nil {
		return -2, fmt.Errorf("ошибка валидации данных: %v", err)
	}

	// Хеширование пароля
	hash, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.MinCost)
	if err != nil {
		return -3, fmt.Errorf("ошибка хеширования пароля: %v", err)
	}

	// Подготовка данных для создания пользователя
	var data models.UserTable
	data.Username = userData.Username
	data.PasswordHash = string(hash)
	data.FullName = fullName

	// Создание пользователя
	id, err := uc.userService.CreateUser(data)
	if err != nil {
		return -4, fmt.Errorf("ошибка при создании пользователя: %v", err)
	}

	return id, nil
}
