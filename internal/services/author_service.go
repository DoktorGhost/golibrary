package services

import (
	"fmt"
	"github.com/DoktorGhost/golibrary/internal/models"
	"github.com/DoktorGhost/golibrary/internal/services/crud"
)

type AuthorService struct {
	repo crud.AutorRepository
}

func NewAuthorService(repo crud.AutorRepository) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) AddAuthor(fullName string) (int, error) {
	id, err := s.repo.CreateAuthor(fullName)
	if err != nil {
		return 0, fmt.Errorf("ошибка создания автора: %v", err)
	}

	return id, nil
}

func (s *AuthorService) DeleteAuthor(id int) error {
	err := s.repo.DeleteAuthor(id)
	if err != nil {
		return fmt.Errorf("ошибка удаления автора: %v", err)
	}
	return nil
}

func (s *AuthorService) UpdateAuthor(author models.AuthorTable) error {
	err := s.repo.UpdateAuthor(author)
	if err != nil {
		return fmt.Errorf("ошибка обновления автора: %v", err)
	}

	return nil
}

func (s *AuthorService) GetAuthorById(id int) (models.AuthorTable, error) {
	author, err := s.repo.GetAuthorByID(id)
	if err != nil {
		return models.AuthorTable{}, err
	}

	return author, nil
}

func (s *AuthorService) GetAllAuthors() (map[int]models.AuthorTable, error) {

	authors, err := s.repo.GetAllAuthors()
	if err != nil {
		return nil, err
	}

	result := make(map[int]models.AuthorTable)

	for _, author := range authors {
		result[author.ID] = author
	}

	return result, nil
}
