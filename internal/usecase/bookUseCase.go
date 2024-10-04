package usecase

import (
	"github.com/DoktorGhost/golibrary/internal/models"
	"github.com/DoktorGhost/golibrary/internal/services"
	"github.com/DoktorGhost/golibrary/pkg/validator"
)

type BookUseCase struct {
	bookService   services.BookService
	authorService services.AuthorService
	rentalService services.RentalService
}

func NewBookUseCase(bookService services.BookService, authorService services.AuthorService, rentalService services.RentalService) *BookUseCase {
	return &BookUseCase{bookService: bookService, authorService: authorService, rentalService: rentalService}
}

func (uc *BookUseCase) AddBook(book models.BookTable) (int, error) {
	//Проверяем наличие автора в таблице
	_, err := uc.authorService.GetAuthorById(book.AuthorID)
	if err != nil {
		return -1, err
	}

	id, err := uc.bookService.AddBook(book)
	if err != nil {
		return -2, err
	}
	//добавляем запись в таблицу Rentals, указываем 0, так как книга свободна
	err = uc.rentalService.CreateRentals(id)
	if err != nil {
		return -3, err
	}
	return id, nil
}

func (uc *BookUseCase) AddAuthor(name, surname, patronymic string) (int, error) {
	fullName, err := validator.Valid(name, surname, patronymic)
	if err != nil {
		return -1, err
	}
	id, err := uc.authorService.AddAuthor(fullName)
	if err != nil {
		return -2, err
	}
	return id, nil
}

func (uc *BookUseCase) GetAllBookWithAuthor() ([]models.Book, error) {
	authors, err := uc.authorService.GetAllAuthors()
	if err != nil {
		return nil, err
	}
	books, err := uc.bookService.GetAllBook()

	var bookList []models.Book

	for _, bookTable := range books {
		var book models.Book
		book.ID = bookTable.ID
		book.Title = bookTable.Title
		book.Author = authors[bookTable.AuthorID]
		bookList = append(bookList, book)
	}
	return bookList, nil
}

func (uc *BookUseCase) GetBookWithAuthor(id int) (models.Book, error) {
	bookTable, err := uc.bookService.GetBook(id)
	if err != nil {
		return models.Book{}, err
	}

	author, err := uc.authorService.GetAuthorById(bookTable.AuthorID)
	if err != nil {
		return models.Book{}, err
	}

	book := models.Book{
		ID:     bookTable.ID,
		Title:  bookTable.Title,
		Author: author,
	}

	return book, nil

}
