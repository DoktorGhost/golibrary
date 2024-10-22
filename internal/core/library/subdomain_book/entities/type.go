package entities

import (
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/repositories/postgres/dao"
)

type Book struct {
	ID     int             `json:"id"`
	Title  string          `json:"title"`
	Author dao.AuthorTable `json:"author"`
}

type Author struct {
	ID    int             `json:"id"`
	Name  string          `json:"name"`
	Books []dao.BookTable `json:"books"`
}

type AuthorRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type BookRequest struct {
	Title    string `json:"title"`
	AuthorID int    `json:"author_id"`
}
