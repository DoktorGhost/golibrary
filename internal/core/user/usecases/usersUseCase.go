package usecases

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/core/user/entities"
	"github.com/DoktorGhost/golibrary/internal/core/user/services"
	"strconv"
)

type UsersUseCase struct {
	userService *services.UserService
}

func NewUsersUseCase(userService *services.UserService) *UsersUseCase {
	return &UsersUseCase{userService: userService}
}

func (uc *UsersUseCase) AddUser(userData entities.RegisterData) (int, error) {
	idStr, err := uc.userService.Register(userData)
	if err != nil {
		return 0, fmt.Errorf("ошибка при создании пользователя: %v", err)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc *UsersUseCase) GetUserByID(id int) (string, error) {
	username, err := uc.userService.GetUserById(id)
	if err != nil {
		return "", err
	}

	return username, nil
}
