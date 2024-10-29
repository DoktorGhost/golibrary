package usecases

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/entities"

	services2 "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/services"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/services"
)

type BookUseCase struct {
	bookService   *services2.BookService
	rentalService *services.RentalService
}

func NewBookUseCase(
	bookService *services2.BookService,
	rentalService *services.RentalService,
) *BookUseCase {
	return &BookUseCase{bookService: bookService, rentalService: rentalService}
}

func (uc *BookUseCase) AddBook(book entities.BookRequest) (int, error) {

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
	id, err := uc.bookService.AddAuthor(name, surname, patronymic)
	if err != nil {
		return 0, fmt.Errorf("ошибка добавления автора: %v", err)
	}
	return id, nil
}

func (uc *BookUseCase) GetAllBookWithAutor() ([]entities.Book, error) {
	books, err := uc.bookService.GetAllBookWithAutor()
	if err != nil {
		return nil, fmt.Errorf("ошибка получения всех авторов: %v", err)
	}

	return books, nil
}

func (uc *BookUseCase) GetBookWithAutor(id int) (entities.Book, error) {
	book, err := uc.bookService.GetBookWithAutor(id)
	if err != nil {
		return entities.Book{}, fmt.Errorf("ошибка получения книги: %v", err)
	}
	return book, nil
}

func (uc *BookUseCase) GetAllAuthorWithBooks() ([]entities.Author, error) {
	authors, err := uc.bookService.GetAllAuthorWithBooks()
	if err != nil {
		return nil, fmt.Errorf("ошибка получения всех авторов: %v", err)
	}

	return authors, nil
}
