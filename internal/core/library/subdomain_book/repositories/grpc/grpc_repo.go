package postgres

import (
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/entities"
	"github.com/DoktorGhost/golibrary/internal/delivery/grpc/client"
)

type BookRepository struct {
	client *client.BookClient
}

func NewBookRepository(client *client.BookClient) *BookRepository {
	return &BookRepository{client: client}
}

func (br *BookRepository) AddBook(book entities.BookRequest) (int, error) {
	return br.client.AddBook(book)
}

func (br *BookRepository) AddAuthor(name, surname, patronymic string) (int, error) {
	return br.client.AddAuthor(name, surname, patronymic)
}

func (br *BookRepository) GetAllBookWithAutor() ([]entities.Book, error) {
	return br.client.GetAllBookWithAutor()
}

func (br *BookRepository) GetBookWithAutor(id int) (entities.Book, error) {
	return br.client.GetBookWithAutor(id)
}

func (br *BookRepository) GetAllAuthorWithBooks() ([]entities.Author, error) {
	return br.client.GetAllAuthorWithBooks()
}
