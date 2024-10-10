package usecases

import (
	"errors"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/entities"
	"github.com/DoktorGhost/golibrary/internal/core/user/repositories/postgres/dao"
	"github.com/DoktorGhost/golibrary/internal/core/user/services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := services.NewMockUserRepository(ctrl)
	userService := services.NewUserService(mockRepo)
	useCase := NewUsersUseCase(userService)

	// 1. Успешный сценарий
	t.Run("CreateUser success", func(t *testing.T) {
		user := entities.RegisterData{
			Username:   "test_user",
			Password:   "password123",
			Name:       "Test",
			Surname:    "Testov",
			Patronymic: "Testovich",
		}

		mockRepo.EXPECT().GetUserByUsername(user.Username).Return(dao.UserTable{}, errors.New("ERROR"))
		mockRepo.EXPECT().CreateUser(gomock.Any()).Return(10, nil)

		userID, err := useCase.AddUser(user)

		assert.NoError(t, err)
		assert.Equal(t, 10, userID)
	})
	// 2. Дублирующийся username
	t.Run("CreateUser success", func(t *testing.T) {
		user := entities.RegisterData{}

		mockRepo.EXPECT().GetUserByUsername(user.Username).Return(dao.UserTable{}, nil)

		userID, err := useCase.AddUser(user)

		assert.Error(t, err)
		assert.Equal(t, 0, userID)
	})
	// 3. Невалидные ФИО
	t.Run("CreateUser success", func(t *testing.T) {
		user := entities.RegisterData{
			Username:   "test_user",
			Password:   "password123",
			Name:       "Test1", //ошибка в имени
			Surname:    "Testov",
			Patronymic: "Testovich",
		}

		mockRepo.EXPECT().GetUserByUsername(user.Username).Return(dao.UserTable{}, errors.New("ERROR"))

		userID, err := useCase.AddUser(user)

		assert.Error(t, err)
		assert.Equal(t, 0, userID)
	})
	// 4. Ошибка создания пользователя
	t.Run("CreateUser success", func(t *testing.T) {
		user := entities.RegisterData{
			Username:   "test_user",
			Password:   "",
			Name:       "Test", //ошибка в имени
			Surname:    "Testov",
			Patronymic: "Testovich",
		}

		mockRepo.EXPECT().GetUserByUsername(user.Username).Return(dao.UserTable{}, errors.New("ERROR"))
		mockRepo.EXPECT().CreateUser(gomock.Any()).Return(0, errors.New("ERROR"))

		userID, err := useCase.AddUser(user)

		assert.Error(t, err)
		assert.Equal(t, 0, userID)
	})

}

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := services.NewMockUserRepository(ctrl)
	userService := services.NewUserService(mockRepo)
	useCase := NewUsersUseCase(userService)

	// 1. Успешный сценарий
	t.Run("GetUserByID success", func(t *testing.T) {
		userID := 10

		mockRepo.EXPECT().GetUserByID(userID).Return(dao.UserTable{ID: 10}, nil)

		user, err := useCase.GetUserByID(userID)

		assert.NoError(t, err)
		assert.Equal(t, 10, user.ID)
	})
	// 2. Сценарий с ошибкой
	t.Run("GetUserByID fail", func(t *testing.T) {

		userID := 10
		mockRepo.EXPECT().GetUserByID(userID).Return(dao.UserTable{}, errors.New("ERROR"))

		user, err := useCase.GetUserByID(userID)

		assert.Error(t, err)
		assert.Equal(t, dao.UserTable{}, user)
	})
}
