package services

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/models"
	"github.com/DoktorGhost/golibrary/internal/services/crud"
)

type BookService struct {
	repo crud.BookRepository
}

func NewBookService(repo crud.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) AddBook(book models.BookTable) (int, error) {
	bookID, err := s.repo.CreateBook(book)
	if err != nil {
		return 0, fmt.Errorf("ошибка создания книги: %v", err)
	}
	return bookID, nil
}

func (s *BookService) GetBook(id int) (models.BookTable, error) {
	book, err := s.repo.GetBookByID(id)
	if err != nil {
		return models.BookTable{}, err
	}
	return book, nil
}

func (s *BookService) DeleteBook(id int) error {
	err := s.repo.DeleteBook(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookService) UpdateBook(book models.BookTable) error {
	err := s.repo.UpdateBook(book)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookService) GetAllBook() ([]models.BookTable, error) {
	books, err := s.repo.GetAllBooks()
	if err != nil {
		return nil, err
	}
	var result []models.BookTable

	for _, bookTable := range books {
		var book models.BookTable
		book.Title = bookTable.Title
		book.AuthorID = bookTable.AuthorID
		result = append(result, book)
	}

	return result, nil
}
