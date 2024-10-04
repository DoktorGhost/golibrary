package crud

import "github.com/DoktorGhost/golibrary/internal/models"

// Repository объединяет все специализированные интерфейсы
type Repository interface {
	UserRepository
	BookRepository
	AutorRepository
	RentalsRepository
}

// UserRepository определяет методы для работы с пользователями
type UserRepository interface {
	CreateUser(user models.UserTable) (int, error)
	GetUserByID(id int) (models.UserTable, error)
	UpdateUser(user models.UserTable) error
	DeleteUser(id int) error
	GetUserByUsername(username string) (models.UserTable, error)
}

// BookRepository определяет методы для работы с книгами
type BookRepository interface {
	CreateBook(book models.BookTable) (int, error)
	GetBookByID(id int) (models.BookTable, error)
	UpdateBook(book models.BookTable) error
	DeleteBook(id int) error
	GetAllBooks() ([]models.BookTable, error)
}

type AutorRepository interface {
	CreateAuthor(name string) (int, error)
	GetAuthorByID(id int) (models.AuthorTable, error)
	UpdateAuthor(author models.AuthorTable) error
	DeleteAuthor(id int) error
	GetAllAuthors() ([]models.AuthorTable, error)
}

type RentalsRepository interface {
	CreateRentalsInfo(userID, bookID int) (int, error)
	GetRentalsInfoByID(id int) (models.RentalsTable, error)
	UpdateRentalsInfo(rentals models.RentalsTable) error
	DeleteRentalsInfo(id int) error
	CreateRentals(bookID int) error
	GetRentalsByID(id int) (int, error)
	UpdateRentals(id, rentals_id int) error
	DeleteRentals(id int) error
}
