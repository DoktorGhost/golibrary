package services

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/metrics"
	"time"

	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/repositories/postgres/dao"
)

//go:generate mockgen -source=$GOFILE -destination=./mock_book.go -package=${GOPACKAGE}
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
	start := time.Now()

	bookID, err := s.repo.CreateBook(book)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("AddBook", duration)

	if err != nil {
		return 0, fmt.Errorf("ошибка создания книги: %v", err)
	}

	return bookID, nil
}

func (s *BookService) GetBook(id int) (dao.BookTable, error) {
	start := time.Now()

	book, err := s.repo.GetBookByID(id)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("GetBook", duration)

	if err != nil {
		return dao.BookTable{}, fmt.Errorf("ошибка получения книги: %v", err)
	}

	return book, nil
}

func (s *BookService) DeleteBook(id int) error {
	start := time.Now()

	err := s.repo.DeleteBook(id)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("DeleteBook", duration)

	if err != nil {
		return fmt.Errorf("ошибка удаления книги: %v", err)
	}

	return nil
}

func (s *BookService) UpdateBook(book dao.BookTable) error {
	start := time.Now()

	err := s.repo.UpdateBook(book)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("UpdateBook", duration)

	if err != nil {
		return fmt.Errorf("ошибка обновления книги: %v", err)
	}

	return nil
}

func (s *BookService) GetAllBook() ([]dao.BookTable, error) {
	start := time.Now()

	books, err := s.repo.GetAllBooks()

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("GetAllBook", duration)

	if err != nil {
		return nil, fmt.Errorf("ошибка получения всех книг: %v", err)
	}

	return books, nil
}
