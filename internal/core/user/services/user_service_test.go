package services

import (
	"errors"
	"github.com/DoktorGhost/golibrary/internal/core/user/repositories/postgres/dao"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	// 1. CreateUser успешный сценарий
	t.Run("CreateUser success", func(t *testing.T) {
		user := dao.UserTable{
			ID:           1,
			Username:     "test_user_1",
			PasswordHash: "hash_password",
			FullName:     "Test Test Test",
		}

		mockRepo.EXPECT().CreateUser(user).Return(1, nil)
		userID, err := userService.CreateUser(user)

		assert.NoError(t, err)
		assert.Equal(t, 1, userID)
	})
	// 2. CreateUser сценарий с ошибкой
	t.Run("CreateUser fail", func(t *testing.T) {
		user := dao.UserTable{}
		mockRepo.EXPECT().CreateUser(user).Return(0, errors.New("ОШИБКА"))

		userID, err := userService.CreateUser(user)
		assert.NotNil(t, err)
		assert.Equal(t, 0, userID)
	})

	// 3. GetUserById успешный сценарий
	t.Run("GetUserById success", func(t *testing.T) {
		userID := 10
		userExpect := dao.UserTable{
			ID:           10,
			Username:     "test_user_1",
			PasswordHash: "hash_password",
			FullName:     "Test Test Test",
		}

		mockRepo.EXPECT().GetUserByID(10).Return(userExpect, nil)

		user, err := userService.GetUserById(userID)
		assert.NoError(t, err)
		assert.Equal(t, userExpect, user)
	})
	// 4. GetUserById сценарий с ошибкой
	t.Run("GetUserById fail", func(t *testing.T) {
		userID := 10
		userExpect := dao.UserTable{}
		mockRepo.EXPECT().GetUserByID(userID).Return(dao.UserTable{}, errors.New("ОШИБКА"))

		user, err := userService.GetUserById(userID)
		assert.NotNil(t, err)
		assert.Equal(t, userExpect, user)
	})
	// 5. GetUserByUsername успешный сценарий
	t.Run("GetUserByUsername success", func(t *testing.T) {
		username := "test_user_1"

		userExpect := dao.UserTable{
			ID:           10,
			Username:     "test_user_1",
			PasswordHash: "hash_password",
			FullName:     "Test Test Test",
		}

		mockRepo.EXPECT().GetUserByUsername(username).Return(userExpect, nil)

		user, err := userService.GetUserByUsername(username)
		assert.NoError(t, err)
		assert.Equal(t, userExpect, user)
	})
	// 6. GetUserByUsername сценарий с ошибкой
	t.Run("GetUserByUsername fail", func(t *testing.T) {
		username := "test_user_1"
		userExpect := dao.UserTable{}
		mockRepo.EXPECT().GetUserByUsername(username).Return(dao.UserTable{}, errors.New("ОШИБКА"))

		user, err := userService.GetUserByUsername(username)
		assert.NotNil(t, err)
		assert.Equal(t, userExpect, user)
	})
	// 7. DeleteUser успешный сценарий
	t.Run("DeleteUser success", func(t *testing.T) {
		userID := 10

		mockRepo.EXPECT().DeleteUser(userID).Return(nil)

		err := userService.DeleteUser(userID)
		assert.NoError(t, err)
	})
	// 8. DeleteUser сценарий с ошибкой
	t.Run("DeleteUser fail", func(t *testing.T) {
		userID := 10

		mockRepo.EXPECT().DeleteUser(userID).Return(errors.New("ОШИБКА"))

		err := userService.DeleteUser(userID)
		assert.NotNil(t, err)
	})
	// 9. UpdateUser успешный сценарий
	t.Run("UpdateUser success", func(t *testing.T) {
		user := dao.UserTable{}

		mockRepo.EXPECT().UpdateUser(user).Return(nil)

		err := userService.UpdateUser(user)
		assert.NoError(t, err)
	})
	// 10. UpdateUser сценарий с ошибкой
	t.Run("UpdateUser fail", func(t *testing.T) {
		user := dao.UserTable{}

		mockRepo.EXPECT().UpdateUser(user).Return(errors.New("ОШИБКА"))

		err := userService.UpdateUser(user)
		assert.NotNil(t, err)
	})

}
