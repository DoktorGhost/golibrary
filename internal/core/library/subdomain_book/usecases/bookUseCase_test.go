package usecases_test

import (
	"errors"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/entities"
	dao2 "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/repositories/postgres/dao"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/usecases"

	service2 "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/services"
	service3 "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//моки
	mockRepoRental := service3.NewMockRentalsRepository(ctrl)
	mockRepoBook := service2.NewMockBookRepository(ctrl)
	mockRepoAuthor := service2.NewMockAuthorRepository(ctrl)

	//сервисы
	rentalService := service3.NewRentalService(mockRepoRental)
	authorService := service2.NewAuthorService(mockRepoAuthor)
	bookService := service2.NewBookService(mockRepoBook)

	//юзкейсы
	bookUseCase := usecases.NewBookUseCase(bookService, authorService, rentalService)

	// 1. AddBook. GetAuthorById fail
	t.Run("GetUserById fail", func(t *testing.T) {
		mockRepoAuthor.EXPECT().GetAuthorByID(gomock.Any()).Return(dao2.AuthorTable{}, errors.New("Error"))

		rentalID, err := bookUseCase.AddBook(dao2.BookTable{})

		assert.Error(t, err)
		assert.Equal(t, 0, rentalID)
	})

	// 2. AddBook. AddBook fail
	t.Run("AddBook fail", func(t *testing.T) {
		mockRepoAuthor.EXPECT().GetAuthorByID(gomock.Any()).Return(dao2.AuthorTable{}, nil)
		mockRepoBook.EXPECT().CreateBook(gomock.Any()).Return(0, errors.New("Error"))

		rentalID, err := bookUseCase.AddBook(dao2.BookTable{})

		assert.Error(t, err)
		assert.Equal(t, 0, rentalID)
	})

	// 3. AddBook. CreateRentals fail
	t.Run("CreateRentals fail", func(t *testing.T) {
		mockRepoAuthor.EXPECT().GetAuthorByID(gomock.Any()).Return(dao2.AuthorTable{}, nil)
		mockRepoBook.EXPECT().CreateBook(gomock.Any()).Return(1, nil)
		mockRepoRental.EXPECT().CreateRentals(1).Return(errors.New("Error"))

		rentalID, err := bookUseCase.AddBook(dao2.BookTable{})

		assert.Error(t, err)
		assert.Equal(t, 0, rentalID)
	})

	// 4. AddBook success
	t.Run("AddBook success", func(t *testing.T) {
		mockRepoAuthor.EXPECT().GetAuthorByID(gomock.Any()).Return(dao2.AuthorTable{}, nil)
		mockRepoBook.EXPECT().CreateBook(gomock.Any()).Return(1, nil)
		mockRepoRental.EXPECT().CreateRentals(1).Return(nil)

		rentalID, err := bookUseCase.AddBook(dao2.BookTable{})

		assert.NoError(t, err)
		assert.Equal(t, 1, rentalID)
	})

	// 5. AddAuthor. Not valid name
	t.Run("Not valid name", func(t *testing.T) {
		rentalID, err := bookUseCase.AddAuthor("Test1", "Test", "Test")

		assert.Error(t, err)
		assert.Equal(t, 0, rentalID)
	})

	// 6. AddAuthor CreateAuthor fail
	t.Run("AddBook fail", func(t *testing.T) {
		mockRepoAuthor.EXPECT().CreateAuthor(gomock.Any()).Return(0, errors.New("Error"))

		rentalID, err := bookUseCase.AddAuthor("Test", "Test", "Test")

		assert.Error(t, err)
		assert.Equal(t, 0, rentalID)
	})

	// 7. AddAuthor success
	t.Run("AddBook fail", func(t *testing.T) {
		mockRepoAuthor.EXPECT().CreateAuthor(gomock.Any()).Return(1, nil)

		rentalID, err := bookUseCase.AddAuthor("Test", "Test", "Test")

		assert.NoError(t, err)
		assert.Equal(t, 1, rentalID)
	})

	// 8. GetAllBookWithAuthor. GetAllAuthors fail
	t.Run("GetAllAuthors fail", func(t *testing.T) {
		mockRepoAuthor.EXPECT().GetAllAuthors().Return(nil, errors.New("Error"))
		var result []entities.Book
		bookList, err := bookUseCase.GetAllBookWithAuthor()

		assert.Error(t, err)
		assert.Equal(t, result, bookList)
	})

	// 9. GetAllBookWithAuthor. GetAllBook fail
	t.Run("GetAllBook fail", func(t *testing.T) {
		var result []entities.Book

		mockRepoAuthor.EXPECT().GetAllAuthors().Return([]dao2.AuthorTable{}, nil)
		mockRepoBook.EXPECT().GetAllBooks().Return(nil, errors.New("Error"))

		bookList, err := bookUseCase.GetAllBookWithAuthor()

		assert.Error(t, err)
		assert.Equal(t, result, bookList)
	})

	// 10. GetAllBookWithAuthor success
	t.Run("GetAllBook success", func(t *testing.T) {
		var result []entities.Book

		mockRepoAuthor.EXPECT().GetAllAuthors().Return([]dao2.AuthorTable{}, nil)
		mockRepoBook.EXPECT().GetAllBooks().Return([]dao2.BookTable{}, nil)

		bookList, err := bookUseCase.GetAllBookWithAuthor()

		assert.NoError(t, err)
		assert.Equal(t, result, bookList)
	})

	// 11. GetBookWithAuthor. GetBook fail
	t.Run("GetBook fail", func(t *testing.T) {

		mockRepoBook.EXPECT().GetBookByID(gomock.Any()).Return(dao2.BookTable{}, errors.New("Error"))

		_, err := bookUseCase.GetBookWithAuthor(1)

		assert.Error(t, err)
	})

	// 12. GetBookWithAuthor. GetAuthorById fail
	t.Run("GetAuthorById fail", func(t *testing.T) {

		mockRepoBook.EXPECT().GetBookByID(gomock.Any()).Return(dao2.BookTable{}, nil)
		mockRepoAuthor.EXPECT().GetAuthorByID(gomock.Any()).Return(dao2.AuthorTable{}, errors.New("Error"))

		_, err := bookUseCase.GetBookWithAuthor(1)

		assert.Error(t, err)
	})

	// 13. GetBookWithAuthor success
	t.Run("GetBookWithAuthor success", func(t *testing.T) {

		mockRepoBook.EXPECT().GetBookByID(gomock.Any()).Return(dao2.BookTable{}, nil)
		mockRepoAuthor.EXPECT().GetAuthorByID(gomock.Any()).Return(dao2.AuthorTable{}, nil)

		_, err := bookUseCase.GetBookWithAuthor(1)

		assert.NoError(t, err)
	})

	// 14. GetAllAuthorWithBooks. GetAllAuthors fail
	t.Run("GetAllAuthors fail", func(t *testing.T) {

		mockRepoAuthor.EXPECT().GetAllAuthors().Return([]dao2.AuthorTable{}, errors.New("Error"))
		_, err := bookUseCase.GetAllAuthorWithBooks()

		assert.Error(t, err)
	})

	// 15. GetAllAuthorWithBooks. GetAllBook fail
	t.Run("GetAllBook fail", func(t *testing.T) {

		mockRepoAuthor.EXPECT().GetAllAuthors().Return([]dao2.AuthorTable{}, nil)
		mockRepoBook.EXPECT().GetAllBooks().Return([]dao2.BookTable{}, errors.New("Error"))

		_, err := bookUseCase.GetAllAuthorWithBooks()

		assert.Error(t, err)
	})

	// 16. GetAllAuthorWithBooks success
	t.Run("GetAllAuthorWithBooks success", func(t *testing.T) {

		mockRepoAuthor.EXPECT().GetAllAuthors().Return([]dao2.AuthorTable{}, nil)
		mockRepoBook.EXPECT().GetAllBooks().Return([]dao2.BookTable{}, nil)

		_, err := bookUseCase.GetAllAuthorWithBooks()

		assert.NoError(t, err)
	})
}
