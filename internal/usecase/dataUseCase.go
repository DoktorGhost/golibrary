package usecase

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/models"
	"github.com/DoktorGhost/golibrary/internal/randomData"
	"github.com/brianvoe/gofakeit/v6"
	"math/rand"
)

type DataUseCase struct {
	BookUseCase
	UsersUseCase
}

func NewDataUseCase(bookUseCase BookUseCase, usersUseCase UsersUseCase) *DataUseCase {
	return &DataUseCase{bookUseCase, usersUseCase}
}

func (uc *DataUseCase) AddLibrary() {
	authors, err := uc.authorService.GetAllAuthors()
	if err != nil {
		return
	}
	//добавляем авторов
	if len(authors) < 1 {
		for i := 0; i < 10; i++ {
			name, surname, patronymic := randomData.GenerateName()
			_, err = uc.AddAuthor(name, surname, patronymic)
			if err != nil {
				fmt.Println("Ошибка добавления автора:", err)
			}
		}

	}

	books, err := uc.bookService.GetAllBook()
	if err != nil {
		return
	}
	//добавляем книги
	if len(books) < 1 {
		for i := 0; i < 100; i++ {
			var book models.BookTable
			book.AuthorID = rand.Intn(10) + 1
			book.Title = randomData.GenerateTitleBook()
			id, err := uc.AddBook(book)
			if err != nil {
				fmt.Println("Ошибка добавления книги:", id, err)
			}
		}
	}

	//добавляем пользователей
	_, err = uc.UsersUseCase.userService.GetUserById(50)
	if err != nil {
		for i := 0; i < 60; i++ {
			user := models.RegisterData{
				Username:   gofakeit.Username(),
				Password:   gofakeit.Password(true, false, false, false, false, 7),
				Name:       gofakeit.LastName(),
				Surname:    gofakeit.LastName(),
				Patronymic: gofakeit.LastName(),
			}
			id, err := uc.UsersUseCase.AddUser(user)
			if err != nil {
				fmt.Println(id, err)
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

}
