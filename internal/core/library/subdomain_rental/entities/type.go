package entities

import "github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/entities"

type UserWithRentedBooks struct {
	ID          int             `json:"id"`
	Username    string          `json:"username"`
	RentedBooks []entities.Book `json:"rented_books"`
}
