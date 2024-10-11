package usecases

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/entities"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/repositories/postgres/dao"
	services2 "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/services"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/services"
	"github.com/DoktorGhost/golibrary/pkg/validator"
)

type BookUseCase struct {
	bookService   *services2.BookService
	authorService *services2.AuthorService
	rentalService *services.RentalService
}

func NewBookUseCase(
	bookService *services2.BookService,
	authorService *services2.AuthorService,
	rentalService *services.RentalService,
) *BookUseCase {
	return &BookUseCase{bookService: bookService, authorService: authorService, rentalService: rentalService}
}

func (uc *BookUseCase) AddBook(book dao.BookTable) (int, error) {
	//Проверяем наличие автора в таблице
	_, err := uc.authorService.GetAuthorById(book.AuthorID)
	if err != nil {
		return 0, fmt.Errorf("данного автора id=%d нет в таблице: %v", book.AuthorID, err)
	}

	//Добавляем книгу
	bookID, err := uc.bookService.AddBook(book)
	if err != nil {
		return 0, fmt.Errorf("ошибка добавления книги: %v", err)
	}

	//Добавляем запись в таблицу Rentals (по дефотлту она будет свободна)
	err = uc.rentalService.CreateRentals(bookID)
	if err != nil {
		return 0, fmt.Errorf("ошибка добавления записи в Rentals: %v", err)
	}

	return bookID, nil
}

func (uc *BookUseCase) AddAuthor(name, surname, patronymic string) (int, error) {
	fullName, err := validator.Valid(name, surname, patronymic)
	if err != nil {
		return 0, fmt.Errorf("ошибка валидации ФИО: %v", err)
	}
	id, err := uc.authorService.AddAuthor(fullName)
	if err != nil {
		return 0, fmt.Errorf("ошибка добавления автора: %v", err)
	}
	return id, nil
}

func (uc *BookUseCase) GetAllBookWithAuthor() ([]entities.Book, error) {
	authors, err := uc.authorService.GetAllAuthors()
	if err != nil {
		return nil, fmt.Errorf("ошибка получения всех авторов: %v", err)
	}
	books, err := uc.bookService.GetAllBook()
	if err != nil {
		return nil, fmt.Errorf("ошибка получения всех книг: %v", err)
	}

	var bookList []entities.Book

	for _, bookTable := range books {
		var book entities.Book
		book.ID = bookTable.ID
		book.Title = bookTable.Title
		book.Author = authors[bookTable.AuthorID]
		bookList = append(bookList, book)
	}
	return bookList, nil
}

func (uc *BookUseCase) GetBookWithAuthor(id int) (entities.Book, error) {
	bookTable, err := uc.bookService.GetBook(id)
	if err != nil {
		return entities.Book{}, fmt.Errorf("ошибка получения книги: %v", err)
	}

	author, err := uc.authorService.GetAuthorById(bookTable.AuthorID)
	if err != nil {
		return entities.Book{}, fmt.Errorf("ошибка получения автора: %v", err)
	}

	book := entities.Book{
		ID:     bookTable.ID,
		Title:  bookTable.Title,
		Author: author,
	}

	return book, nil
}

func (uc *BookUseCase) GetAllAuthorWithBooks() ([]entities.Author, error) {
	authors, err := uc.authorService.GetAllAuthors()
	if err != nil {
		return nil, fmt.Errorf("ошибка получения всех авторов: %v", err)
	}
	books, err := uc.bookService.GetAllBook()
	if err != nil {
		return nil, fmt.Errorf("ошибка получения всех книг: %v", err)
	}

	var authorList []entities.Author
	authorBooks := make(map[int][]dao.BookTable)

	for _, bookTable := range books {
		authorBooks[bookTable.AuthorID] = append(authorBooks[bookTable.AuthorID], bookTable)
	}

	for key, value := range authors {
		var author entities.Author
		author.ID = key
		author.Name = value.Name
		author.Books = authorBooks[key]
		authorList = append(authorList, author)
	}
	return authorList, nil
}