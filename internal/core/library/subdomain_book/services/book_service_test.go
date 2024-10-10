package services

import (
	"errors"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/repositories/postgres/dao"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockBookRepository(ctrl)
	userService := NewBookService(mockRepo)

	// 1. CreateBook успешный сценарий
	t.Run("CreateRentalInfo success", func(t *testing.T) {
		book := dao.BookTable{}
		mockRepo.EXPECT().CreateBook(book).Return(1, nil)

		bookID, err := userService.AddBook(book)

		assert.NoError(t, err)
		assert.Equal(t, 1, bookID)
	})

	// 2. CreateBook fail
	t.Run("CreateRentalInfo fail", func(t *testing.T) {
		book := dao.BookTable{}
		mockRepo.EXPECT().CreateBook(book).Return(0, errors.New("ERROR"))

		bookID, err := userService.AddBook(book)

		assert.Error(t, err)
		assert.Equal(t, 0, bookID)
	})

	// 3. GetBook успешный сценарий
	t.Run("GetBook success", func(t *testing.T) {
		id := 1
		mockRepo.EXPECT().GetBookByID(id).Return(dao.BookTable{}, nil)

		bookID, err := userService.GetBook(id)

		assert.NoError(t, err)
		assert.Equal(t, dao.BookTable{}, bookID)
	})

	// 4. GetBook fail
	t.Run("GetBook fail", func(t *testing.T) {
		id := 1
		mockRepo.EXPECT().GetBookByID(id).Return(dao.BookTable{}, errors.New("ERROR"))

		bookID, err := userService.GetBook(id)

		assert.Error(t, err)
		assert.Equal(t, dao.BookTable{}, bookID)
	})

	// 5. DeleteBook успешный сценарий
	t.Run("DeleteBook success", func(t *testing.T) {
		id := 1
		mockRepo.EXPECT().DeleteBook(id).Return(nil)

		err := userService.DeleteBook(id)

		assert.NoError(t, err)

	})

	// 6. DeleteBook fail
	t.Run("DeleteBook fail", func(t *testing.T) {
		id := 1
		mockRepo.EXPECT().DeleteBook(id).Return(errors.New("ERROR"))

		err := userService.DeleteBook(id)

		assert.Error(t, err)

	})

	// 7. UpdateBook успешный сценарий
	t.Run("UpdateBook success", func(t *testing.T) {
		book := dao.BookTable{}
		mockRepo.EXPECT().UpdateBook(book).Return(nil)

		err := userService.UpdateBook(book)

		assert.NoError(t, err)

	})

	// 8. UpdateBook fail
	t.Run("UpdateBook fail", func(t *testing.T) {
		book := dao.BookTable{}
		mockRepo.EXPECT().UpdateBook(book).Return(errors.New("ERROR"))

		err := userService.UpdateBook(book)

		assert.Error(t, err)

	})
	// 9. GetAllBooks успешный сценарий
	t.Run("GetAllBooks success", func(t *testing.T) {
		books := []dao.BookTable{}
		mockRepo.EXPECT().GetAllBooks().Return(books, nil)

		result, err := userService.GetAllBook()

		assert.NoError(t, err)
		assert.Equal(t, books, result)
	})

	// 10. GetAllBooks fail
	t.Run("GetAllBooks fail", func(t *testing.T) {
		var books []dao.BookTable
		mockRepo.EXPECT().GetAllBooks().Return(books, errors.New("ERROR"))

		result, err := userService.GetAllBook()

		assert.Error(t, err)
		assert.Equal(t, books, result)
	})
}
