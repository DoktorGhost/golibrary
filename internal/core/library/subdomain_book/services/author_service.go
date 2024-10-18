package services

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/metrics"
	"time"

	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/repositories/postgres/dao"
)

//go:generate mockgen -source=$GOFILE -destination=./mock_author.go -package=${GOPACKAGE}
type AuthorRepository interface {
	CreateAuthor(name string) (int, error)
	GetAuthorByID(id int) (dao.AuthorTable, error)
	UpdateAuthor(author dao.AuthorTable) error
	DeleteAuthor(id int) error
	GetAllAuthors() ([]dao.AuthorTable, error)
}

type AuthorService struct {
	repo AuthorRepository
}

func NewAuthorService(repo AuthorRepository) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) AddAuthor(fullName string) (int, error) {
	start := time.Now()

	id, err := s.repo.CreateAuthor(fullName)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("AddAuthor", duration)

	if err != nil {
		return 0, fmt.Errorf("ошибка создания автора: %v", err)
	}

	return id, nil
}

func (s *AuthorService) DeleteAuthor(id int) error {
	start := time.Now()

	err := s.repo.DeleteAuthor(id)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("DeleteAuthor", duration)

	if err != nil {
		return fmt.Errorf("ошибка удаления автора: %v", err)
	}

	return nil
}

func (s *AuthorService) UpdateAuthor(author dao.AuthorTable) error {
	start := time.Now()

	err := s.repo.UpdateAuthor(author)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("UpdateAuthor", duration)

	if err != nil {
		return fmt.Errorf("ошибка обновления автора: %v", err)
	}

	return nil
}

func (s *AuthorService) GetAuthorById(id int) (dao.AuthorTable, error) {
	start := time.Now()

	author, err := s.repo.GetAuthorByID(id)

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("GetAuthorById", duration)

	if err != nil {
		return dao.AuthorTable{}, fmt.Errorf("ошибка получения автора: %v", err)
	}

	return author, nil
}

func (s *AuthorService) GetAllAuthors() (map[int]dao.AuthorTable, error) {
	start := time.Now()

	authors, err := s.repo.GetAllAuthors()

	duration := time.Since(start).Seconds()
	metrics.TrackDBDuration("GetAllAuthors", duration)

	if err != nil {
		return nil, fmt.Errorf("ошибка получения всех авторов: %v", err)
	}

	result := make(map[int]dao.AuthorTable)

	for _, author := range authors {
		result[author.ID] = author
	}

	return result, nil
}
