package usecases

import (
	"fmt"
	subdomainBook "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/entities"
	domainBook "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/services"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/usecases"
	domainRental "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/services"
	"github.com/DoktorGhost/golibrary/internal/core/user/entities"
	"github.com/DoktorGhost/golibrary/pkg/randomData"
	"github.com/brianvoe/gofakeit/v6"
	"math/rand"
)

type DataUseCase struct {
	usecases.BookUseCase
	usersUseCase *UsersUseCase

	bookService   *domainBook.BookService
	rentalService *domainRental.RentalService
}

func NewDataUseCase(
	bookService *domainBook.BookService,
	rentalService *domainRental.RentalService,
	usersUseCase *UsersUseCase,
) *DataUseCase {
	return &DataUseCase{
		BookUseCase:   *usecases.NewBookUseCase(bookService, rentalService),
		usersUseCase:  usersUseCase,
		bookService:   bookService,
		rentalService: rentalService,
	}
}

func (uc *DataUseCase) AddLibrary() error {
	authors, err := uc.bookService.GetAllAuthorWithBooks()
	if err != nil {
		return err
	}
	//добавляем авторов
	if len(authors) < 10 {
		for i := 0; i < 10; i++ {
			name, surname, patronymic := randomData.GenerateName()
			_, err = uc.AddAuthor(name, surname, patronymic)
			if err != nil {
				return fmt.Errorf("ошибка добавления автора: %v", err)
			}
		}

	}

	books, err := uc.bookService.GetAllBookWithAutor()
	if err != nil {
		return err
	}
	//добавляем книги
	if len(books) < 100 {
		for i := 0; i < 100; i++ {
			var book subdomainBook.BookRequest
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
