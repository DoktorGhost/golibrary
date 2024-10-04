package usecases

import (
	"fmt"
	"math/rand"

	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/entities"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/repositories/postgres/dao"
	domainBook "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/services"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/usecases"
	"github.com/DoktorGhost/golibrary/pkg/randomData"
	"github.com/brianvoe/gofakeit/v6"
)

type DataUseCase struct {
	usecases.BookUseCase
	usersUseCase *UsersUseCase

	authorService *domainBook.AuthorService
	bookService   *domainBook.BookService
}

func NewDataUseCase(
	bookService *domainBook.BookService,
	authorService *domainBook.AuthorService,
	usersUseCase *UsersUseCase,
) *DataUseCase {
	return &DataUseCase{
		BookUseCase:   *usecases.NewBookUseCase(bookService, authorService, nil),
		usersUseCase:  usersUseCase,
		authorService: authorService,
		bookService:   bookService,
	}
}

func (uc *DataUseCase) AddLibrary() error {
	authors, err := uc.authorService.GetAllAuthors()
	if err != nil {
		return nil
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
		return nil
	}
	//добавляем книги
	if len(books) < 1 {
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

	//вывести все книги с авторами
	booksWithAuthor, err := uc.GetAllBookWithAuthor()
	if err != nil {
		fmt.Println(err)
	}
	for _, book := range booksWithAuthor {
		fmt.Println(book)
	}

	return nil
}
