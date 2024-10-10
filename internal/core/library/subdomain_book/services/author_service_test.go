package services

import (
	"errors"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/repositories/postgres/dao"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthorService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockAuthorRepository(ctrl)
	userService := NewAuthorService(mockRepo)

	// 1. AddAuthor успешный сценарий
	t.Run("AddAuthor success", func(t *testing.T) {
		fullName := "Test Testov Testovich"
		mockRepo.EXPECT().CreateAuthor(fullName).Return(1, nil)

		bookID, err := userService.AddAuthor(fullName)

		assert.NoError(t, err)
		assert.Equal(t, 1, bookID)
	})

	// 2. AddAuthor  fail
	t.Run("AddAuthor fail", func(t *testing.T) {
		fullName := "Test Testov Testovich"
		mockRepo.EXPECT().CreateAuthor(fullName).Return(0, errors.New("ERROR"))

		bookID, err := userService.AddAuthor(fullName)

		assert.Error(t, err)
		assert.Equal(t, 0, bookID)
	})

	// 3. DeleteAuthor успешный сценарий
	t.Run("DeleteAuthor success", func(t *testing.T) {
		id := 1
		mockRepo.EXPECT().DeleteAuthor(id).Return(nil)

		err := userService.DeleteAuthor(id)

		assert.NoError(t, err)
	})

	// 4. DeleteAuthor  fail
	t.Run("DeleteAuthor fail", func(t *testing.T) {
		id := 1
		mockRepo.EXPECT().DeleteAuthor(id).Return(errors.New("ERROR"))

		err := userService.DeleteAuthor(id)

		assert.Error(t, err)

	})

	// 5. UpdateAuthor успешный сценарий
	t.Run("UpdateAuthor success", func(t *testing.T) {
		author := dao.AuthorTable{}
		mockRepo.EXPECT().UpdateAuthor(author).Return(nil)

		err := userService.UpdateAuthor(author)

		assert.NoError(t, err)
	})

	// 6. UpdateAuthor  fail
	t.Run("UpdateAuthor fail", func(t *testing.T) {
		author := dao.AuthorTable{}
		mockRepo.EXPECT().UpdateAuthor(author).Return(errors.New("ERROR"))

		err := userService.UpdateAuthor(author)

		assert.Error(t, err)

	})

	// 7. GetAuthorById успешный сценарий
	t.Run("GetAuthorById success", func(t *testing.T) {
		id := 1
		mockRepo.EXPECT().GetAuthorByID(id).Return(dao.AuthorTable{}, nil)

		author, err := userService.GetAuthorById(id)

		assert.NoError(t, err)
		assert.Equal(t, dao.AuthorTable{}, author)
	})

	// 8. GetAuthorById  fail
	t.Run("GetAuthorById fail", func(t *testing.T) {
		id := 1
		mockRepo.EXPECT().GetAuthorByID(id).Return(dao.AuthorTable{}, errors.New("ERROR"))

		author, err := userService.GetAuthorById(id)

		assert.Error(t, err)
		assert.Equal(t, dao.AuthorTable{}, author)

	})
	// 9. GetAllAuthors успешный сценарий
	t.Run("GetAllAuthors success", func(t *testing.T) {
		maps := make(map[int]dao.AuthorTable)
		mockRepo.EXPECT().GetAllAuthors().Return([]dao.AuthorTable{}, nil)

		authors, err := userService.GetAllAuthors()

		assert.NoError(t, err)
		assert.Equal(t, maps, authors)
	})

	// 10. GetAllAuthors  fail
	t.Run("GetAllAuthors fail", func(t *testing.T) {
		var maps map[int]dao.AuthorTable
		mockRepo.EXPECT().GetAllAuthors().Return([]dao.AuthorTable{}, errors.New("ERROR"))

		authors, err := userService.GetAllAuthors()

		assert.Error(t, err)
		assert.Equal(t, maps, authors)

	})
}
