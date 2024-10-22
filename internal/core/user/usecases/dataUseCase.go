package usecases

import (
	"fmt"
	domainRental "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/services"
	"math/rand"

	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/repositories/postgres/dao"
	domainBook "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/services"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/usecases"
	"github.com/DoktorGhost/golibrary/internal/core/user/entities"
	"github.com/DoktorGhost/golibrary/pkg/randomData"
	"github.com/brianvoe/gofakeit/v6"
)

type DataUseCase struct {
	usecases.BookUseCase
	usersUseCase *UsersUseCase

	authorService *domainBook.AuthorService
	bookService   *domainBook.BookService
	rentalService *domainRental.RentalService
}

func NewDataUseCase(
	bookService *domainBook.BookService,
	rentalService *domainRental.RentalService,
	authorService *domainBook.AuthorService,
	usersUseCase *UsersUseCase,
) *DataUseCase {
	return &DataUseCase{
		BookUseCase:   *usecases.NewBookUseCase(bookService, authorService, rentalService),
		usersUseCase:  usersUseCase,
		authorService: authorService,
		bookService:   bookService,
		rentalService: rentalService,
	}
}

func (uc *DataUseCase) AddLibrary() error {
	authors, err := uc.authorService.GetAllAuthors()
	if err != nil {
		return err
	}
	//добавляем авторов
	if len(authors) < 1 {
		for i := 0; i < 10; i++ {
			name, surname, patronymic := randomData.GenerateName()
			_, err = uc.AddAuthor(name, surname, patronymic)
			if err != nil {
				return fmt.Errorf("ошибка добавления автора: %v", err)
			}
		}

	}

	books, err := uc.bookService.GetAllBook()
	if err != nil {
		return err
	}
	//добавляем книги
	if len(books) < 100 {
		for i := 0; i < 100; i++ {
			var book dao.BookTable
			book.AuthorID = rand.Intn(10) + 1
			book.Title = randomData.GenerateTitleBook()
			id, err := uc.AddBook(book)
			if err != nil {
				return fmt.Errorf("ошибка добавления книги: %d %v", id, err)
			}
		}
	}

	//добавляем пользователей
	_, err = uc.usersUseCase.GetUserByID(50)
	if err != nil {
		for i := 0; i < 60; i++ {
			user := entities.RegisterData{
				Username:   gofakeit.Username(),
				Password:   gofakeit.Password(true, false, false, false, false, 7),
				Name:       gofakeit.LastName(),
				Surname:    gofakeit.LastName(),
				Patronymic: gofakeit.LastName(),
			}
			id, err := uc.usersUseCase.AddUser(user)
			if err != nil {
				return fmt.Errorf("ошибка добавления пользователя: %d %v", id, err)
			}
		}

	}

	return nil
}
