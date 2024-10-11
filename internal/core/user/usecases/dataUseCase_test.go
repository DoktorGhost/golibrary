package usecases

import (
	"errors"
	dao2 "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/repositories/postgres/dao"
	service2 "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/services"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/usecases"
	service3 "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/services"
	"github.com/DoktorGhost/golibrary/internal/core/user/repositories/postgres/dao"
	"github.com/DoktorGhost/golibrary/internal/core/user/services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddLibrary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//моки
	mockRepoUser := services.NewMockUserRepository(ctrl)
	mockRepoBook := service2.NewMockBookRepository(ctrl)
	mockRepoRental := service3.NewMockRentalsRepository(ctrl)
	mockRepoAuthor := service2.NewMockAuthorRepository(ctrl)

	//сервисы
	bookService := service2.NewBookService(mockRepoBook)
	rentalService := service3.NewRentalService(mockRepoRental)
	authorService := service2.NewAuthorService(mockRepoAuthor)
	userService := services.NewUserService(mockRepoUser)

	//бзкейсы
	userUseCase := NewUsersUseCase(userService)
	bookUseCase := usecases.NewBookUseCase(bookService, authorService, rentalService)
	//bookUseCase := usecases.NewBookUseCase(bookService, authorService, rentalService)

	useCase := &DataUseCase{
		*bookUseCase,
		userUseCase,
		authorService,
		bookService,
		rentalService,
	}
	//useCase := NewDataUseCase(bookService, rentalService, authorService, userUseCase)

	// 1. Ошибка  GetAllAuthors
	t.Run("GetAllAuthors fail", func(t *testing.T) {

		//пустая мапа
		//maps := make(map[int]dao2.AuthorTable)

		mockRepoAuthor.EXPECT().GetAllAuthors().Return(nil, errors.New("ERROR"))

		err := useCase.AddLibrary()

		assert.Error(t, err)
	})

	// 2. Ошибка  AddAuthor
	t.Run("AddAuthor fail", func(t *testing.T) {
		// Пустой слайс авторов
		authors := []dao2.AuthorTable{}

		mockRepoAuthor.EXPECT().GetAllAuthors().Return(authors, nil)
		mockRepoAuthor.EXPECT().CreateAuthor(gomock.Any()).Return(0, errors.New("ERROR"))

		err := useCase.AddLibrary()

		assert.Error(t, err)
	})

	// 3. Ошибка GetAllBook
	t.Run("AddAuthor fail", func(t *testing.T) {
		// Пустой слайс авторов
		authors := []dao2.AuthorTable{}

		mockRepoAuthor.EXPECT().GetAllAuthors().Return(authors, nil)
		mockRepoAuthor.EXPECT().CreateAuthor(gomock.Any()).Return(1, nil).Times(10)
		mockRepoBook.EXPECT().GetAllBooks().Return(nil, errors.New("ERROR"))

		err := useCase.AddLibrary()

		assert.Error(t, err)
	})

	// 4. Ошибка  AddBook
	t.Run("AddBook fail", func(t *testing.T) {
		// Пустой слайс авторов
		authors := []dao2.AuthorTable{}
		books := []dao2.BookTable{}

		mockRepoAuthor.EXPECT().GetAllAuthors().Return(authors, nil)
		mockRepoAuthor.EXPECT().CreateAuthor(gomock.Any()).Return(1, nil).Times(10)
		mockRepoBook.EXPECT().GetAllBooks().Return(books, nil)
		mockRepoAuthor.EXPECT().GetAuthorByID(gomock.Any()).Return(dao2.AuthorTable{}, nil)
		mockRepoBook.EXPECT().CreateBook(gomock.Any()).Return(0, errors.New("ERROR"))

		err := useCase.AddLibrary()

		assert.Error(t, err)
	})

	// 4. Ошибка  GetUserByID
	t.Run("AddBook fail", func(t *testing.T) {
		// Пустой слайс авторов
		authors := []dao2.AuthorTable{}
		books := []dao2.BookTable{}

		mockRepoAuthor.EXPECT().GetAllAuthors().Return(authors, nil)
		mockRepoAuthor.EXPECT().CreateAuthor(gomock.Any()).Return(1, nil).Times(10)
		mockRepoBook.EXPECT().GetAllBooks().Return(books, nil)
		mockRepoAuthor.EXPECT().GetAuthorByID(gomock.Any()).Return(dao2.AuthorTable{}, nil).Times(100)
		mockRepoBook.EXPECT().CreateBook(gomock.Any()).Return(1, nil).Times(100)
		mockRepoRental.EXPECT().CreateRentals(1).Return(nil).Times(100)

		mockRepoUser.EXPECT().GetUserByID(gomock.Any()).Return(dao.UserTable{}, nil)

		err := useCase.AddLibrary()

		assert.NoError(t, err)
	})

	// ошибка AddUser
	t.Run("AddBook fail", func(t *testing.T) {
		// Пустой слайс авторов
		authors := []dao2.AuthorTable{}
		books := []dao2.BookTable{}

		mockRepoAuthor.EXPECT().GetAllAuthors().Return(authors, nil)
		mockRepoAuthor.EXPECT().CreateAuthor(gomock.Any()).Return(1, nil).Times(10)
		mockRepoBook.EXPECT().GetAllBooks().Return(books, nil)
		mockRepoAuthor.EXPECT().GetAuthorByID(gomock.Any()).Return(dao2.AuthorTable{}, nil).Times(100)
		mockRepoBook.EXPECT().CreateBook(gomock.Any()).Return(1, nil).Times(100)
		mockRepoRental.EXPECT().CreateRentals(1).Return(nil).Times(100)
		mockRepoUser.EXPECT().GetUserByID(gomock.Any()).Return(dao.UserTable{}, errors.New("ERROR"))

		mockRepoUser.EXPECT().GetUserByUsername(gomock.Any()).Return(dao.UserTable{}, nil)

		err := useCase.AddLibrary()

		assert.Error(t, err)
	})

	// успешный сценарий
	t.Run("AddBook fail", func(t *testing.T) {
		// Пустой слайс авторов
		authors := []dao2.AuthorTable{}
		books := []dao2.BookTable{}

		mockRepoAuthor.EXPECT().GetAllAuthors().Return(authors, nil)
		mockRepoAuthor.EXPECT().CreateAuthor(gomock.Any()).Return(1, nil).Times(10)
		mockRepoBook.EXPECT().GetAllBooks().Return(books, nil)
		mockRepoAuthor.EXPECT().GetAuthorByID(gomock.Any()).Return(dao2.AuthorTable{}, nil).Times(100)
		mockRepoBook.EXPECT().CreateBook(gomock.Any()).Return(1, nil).Times(100)
		mockRepoRental.EXPECT().CreateRentals(1).Return(nil).Times(100)
		mockRepoUser.EXPECT().GetUserByID(gomock.Any()).Return(dao.UserTable{}, errors.New("ERROR"))
		mockRepoUser.EXPECT().GetUserByUsername(gomock.Any()).Return(dao.UserTable{}, errors.New("ERROR")).Times(60)
		mockRepoUser.EXPECT().CreateUser(gomock.Any()).Return(1, nil).Times(100).Times(60)

		err := useCase.AddLibrary()

		assert.NoError(t, err)
	})

	// успешный сценарий, без добавления авторов
	t.Run("AddBook fail", func(t *testing.T) {
		// Пустой слайс авторов
		authors := []dao2.AuthorTable{}
		for i := 0; i < 11; i++ {
			authors = append(authors, dao2.AuthorTable{})
		}
		books := []dao2.BookTable{}

		mockRepoAuthor.EXPECT().GetAllAuthors().Return(authors, nil)

		mockRepoBook.EXPECT().GetAllBooks().Return(books, nil)
		mockRepoAuthor.EXPECT().GetAuthorByID(gomock.Any()).Return(dao2.AuthorTable{}, nil).Times(100)
		mockRepoBook.EXPECT().CreateBook(gomock.Any()).Return(1, nil).Times(100)
		mockRepoRental.EXPECT().CreateRentals(1).Return(nil).Times(100)
		mockRepoUser.EXPECT().GetUserByID(gomock.Any()).Return(dao.UserTable{}, errors.New("ERROR"))
		mockRepoUser.EXPECT().GetUserByUsername(gomock.Any()).Return(dao.UserTable{}, errors.New("ERROR")).Times(60)
		mockRepoUser.EXPECT().CreateUser(gomock.Any()).Return(1, nil).Times(100).Times(60)

		err := useCase.AddLibrary()

		assert.NoError(t, err)
	})

	// успешный сценарий, без добавления книг
	t.Run("AddBook fail", func(t *testing.T) {
		// Пустой слайс авторов
		authors := []dao2.AuthorTable{}

		books := []dao2.BookTable{}
		for i := 0; i < 100; i++ {
			books = append(books, dao2.BookTable{})
		}

		mockRepoAuthor.EXPECT().GetAllAuthors().Return(authors, nil)
		mockRepoAuthor.EXPECT().CreateAuthor(gomock.Any()).Return(1, nil).Times(10)
		mockRepoBook.EXPECT().GetAllBooks().Return(books, nil)

		mockRepoUser.EXPECT().GetUserByID(gomock.Any()).Return(dao.UserTable{}, errors.New("ERROR"))
		mockRepoUser.EXPECT().GetUserByUsername(gomock.Any()).Return(dao.UserTable{}, errors.New("ERROR")).Times(60)
		mockRepoUser.EXPECT().CreateUser(gomock.Any()).Return(1, nil).Times(100).Times(60)

		err := useCase.AddLibrary()

		assert.NoError(t, err)
	})

	// успешный сценарий, без добавления пользователей
	t.Run("AddBook fail", func(t *testing.T) {
		// Пустой слайс авторов
		authors := []dao2.AuthorTable{}
		books := []dao2.BookTable{}

		mockRepoAuthor.EXPECT().GetAllAuthors().Return(authors, nil)
		mockRepoAuthor.EXPECT().CreateAuthor(gomock.Any()).Return(1, nil).Times(10)
		mockRepoBook.EXPECT().GetAllBooks().Return(books, nil)
		mockRepoAuthor.EXPECT().GetAuthorByID(gomock.Any()).Return(dao2.AuthorTable{}, nil).Times(100)
		mockRepoBook.EXPECT().CreateBook(gomock.Any()).Return(1, nil).Times(100)
		mockRepoRental.EXPECT().CreateRentals(1).Return(nil).Times(100)
		mockRepoUser.EXPECT().GetUserByID(gomock.Any()).Return(dao.UserTable{}, nil)

		err := useCase.AddLibrary()

		assert.NoError(t, err)
	})

}
