package services

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/models"
	"github.com/DoktorGhost/golibrary/internal/services/crud"
)

type BookService struct {
	repo crud.Repository
}

func (s *BookService) AddBook(book models.Book) (int, error) {
	// Проверка существования автора
	_, err := s.repo.GetAuthorByID(book.AuthorID)
	if err != nil {
		return 0, fmt.Errorf("автор с ID %d не найден: %v", book.AuthorID, err)
	}

	// Создание книги
	bookID, err := s.repo.CreateBook(book)
	if err != nil {
		return 0, fmt.Errorf("ошибка создания книги: %v", err)
	}

	return bookID, nil

}

func (s *BookService) UpdateBookDetails(book models.Book) error {
	// бизнес-логика
	return s.repo.UpdateBook(book)
}
