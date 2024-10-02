package crud

import "github.com/DoktorGhost/golibrary/internal/models"

// UserRepository определяет методы для работы с пользователями
type UserRepository interface {
	CreateUser(user models.User) (int, error)
	GetUserByID(id int) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id int) error
}

// BookRepository определяет методы для работы с книгами
type BookRepository interface {
	CreateBook(book models.Book) (int, error)
	GetBookByID(id int) (models.Book, error)
	UpdateBook(book models.Book) error
	DeleteBook(id int) error
}

type AutorRepository interface {
	CreateAuthor(book models.Author) (int, error)
	GetAuthorByID(id int) (models.Author, error)
	UpdateAuthor(book models.Author) error
	DeleteAuthor(id int) error
}

// Repository объединяет все специализированные интерфейсы
type Repository interface {
	UserRepository
	BookRepository
	AutorRepository
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}
