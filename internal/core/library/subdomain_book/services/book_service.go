package services

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/entities"
	"github.com/DoktorGhost/golibrary/internal/metrics"
	"time"
)

//go:generate mockgen -source=$GOFILE -destination=./mock_book.go -package=${GOPACKAGE}
type BookRepository interface {
	AddBook(book entities.BookRequest) (int, error)
	AddAuthor(name, surname, patronymic string) (int, error)
	GetAllBookWithAutor() ([]entities.Book, error)
	GetBookWithAutor(id int) (entities.Book, error)
	GetAllAuthorWithBooks() ([]entities.Author, error)
}

type BookService struct {
	repo BookRepository
}

func NewBookService(repo BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) AddBook(book entities.BookRequest) (int, error) {
	start := time.Now()

	bookID, err := s.repo.AddBook(book)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("AddBook", duration)

	if err != nil {
		return 0, fmt.Errorf("ошибка создания книги: %v", err)
	}

	return bookID, nil
}

func (s *BookService) AddAuthor(name, surname, patronymic string) (int, error) {
	start := time.Now()

	authorID, err := s.repo.AddAuthor(name, surname, patronymic)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("AddAuthor", duration)

	if err != nil {
		return 0, fmt.Errorf("ошибка создания автора: %v", err)
	}

	return authorID, nil
}

func (s *BookService) GetAllBookWithAutor() ([]entities.Book, error) {
	start := time.Now()

	books, err := s.repo.GetAllBookWithAutor()

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("GetAllBookWithAutor", duration)

	if err != nil {
		return nil, fmt.Errorf("ошибка получения всех книг: %v", err)
	}

	return books, nil
}

func (s *BookService) GetBookWithAutor(id int) (entities.Book, error) {
	start := time.Now()

	book, err := s.repo.GetBookWithAutor(id)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("GetBookWithAutor", duration)

	if err != nil {
		return entities.Book{}, fmt.Errorf("ошибка получения книги: %v", err)
	}

	return book, nil
}

func (s *BookService) GetAllAuthorWithBooks() ([]entities.Author, error) {
	start := time.Now()

	authors, err := s.repo.GetAllAuthorWithBooks()

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("GetAllAuthorWithBooks", duration)

	if err != nil {
		return nil, fmt.Errorf("ошибка получения книги: %v", err)
	}

	return authors, nil
}
