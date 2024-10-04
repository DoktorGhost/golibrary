package services

import (
	"fmt"

	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/repositories/postgres/dao"
)

type BookRepository interface {
	CreateBook(book dao.BookTable) (int, error)
	GetBookByID(id int) (dao.BookTable, error)
	UpdateBook(book dao.BookTable) error
	DeleteBook(id int) error
	GetAllBooks() ([]dao.BookTable, error)
}

type BookService struct {
	repo BookRepository
}

func NewBookService(repo BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) AddBook(book dao.BookTable) (int, error) {
	bookID, err := s.repo.CreateBook(book)
	if err != nil {
		return 0, fmt.Errorf("ошибка создания книги: %v", err)
	}
	return bookID, nil
}

func (s *BookService) GetBook(id int) (dao.BookTable, error) {
	book, err := s.repo.GetBookByID(id)
	if err != nil {
		return dao.BookTable{}, fmt.Errorf("ошибка получения книги: %v", err)
	}
	return book, nil
}

func (s *BookService) DeleteBook(id int) error {
	err := s.repo.DeleteBook(id)
	if err != nil {
		return fmt.Errorf("ошибка удаления книги: %v", err)
	}
	return nil
}

func (s *BookService) UpdateBook(book dao.BookTable) error {
	err := s.repo.UpdateBook(book)
	if err != nil {
		return fmt.Errorf("ошибка обновления книги: %v", err)
	}
	return nil
}

func (s *BookService) GetAllBook() ([]dao.BookTable, error) {
	books, err := s.repo.GetAllBooks()
	if err != nil {
		return nil, fmt.Errorf("ошибка получения всех книг: %v", err)
	}
	return books, nil
}
