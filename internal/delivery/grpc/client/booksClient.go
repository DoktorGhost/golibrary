package client

import (
	"context"
	proto "github.com/DoktorGhost/external-api/src/go/pkg/grpc/clients/api/grpc/protobuf/books_v1"
	"github.com/DoktorGhost/golibrary/config"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/entities"
	"google.golang.org/grpc"
	"log"
	"time"
)

type BookClient struct {
	proto.BooksServiceClient
}

func InitBookClient() (*BookClient, *grpc.ClientConn) {
	// Подключаемся к gRPC-серверу USER
	conn, err := grpc.Dial(config.LoadConfig().GrpcConfig.BookHost+":"+config.LoadConfig().GrpcConfig.BookPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// Создаем gRPC-клиента для сервиса USER
	booksClient := proto.NewBooksServiceClient(conn)

	// Создаем сервис User, который будет использовать этот клиент
	bookService := &BookClient{booksClient}

	log.Println("Connected to Book service port:", config.LoadConfig().GrpcConfig.BookPort)
	return bookService, conn
}

func (a *BookClient) AddBook(book entities.BookRequest) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Вызываем метод AddBook в сервисе Books

	resp, err := a.BooksServiceClient.AddBook(ctx, &proto.AddBookRequest{Title: book.Title, AuthorId: int64(book.AuthorID)})
	if err != nil {
		return 0, err
	}
	return int(resp.Id), nil
}

func (a *BookClient) AddAuthor(name, surname, patronymic string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Вызываем метод AddAuthor в сервисе Books
	resp, err := a.BooksServiceClient.AddAuthor(ctx, &proto.AddAuthorRequest{Name: name, Surname: surname, Patronymic: patronymic})
	if err != nil {
		return 0, err
	}
	return int(resp.Id), nil
}

func (a *BookClient) GetAllBookWithAutor() ([]entities.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Вызываем метод GetAllBookWithAuthor в сервисе Books
	resp, err := a.BooksServiceClient.GetAllBookWithAuthor(ctx, &proto.Empty{})
	if err != nil {
		return nil, err
	}

	// Создаем слайс для хранения книг
	books := make([]entities.Book, len(resp.Books))

	// Итерируем по результатам и конвертируем их в entities.Book
	for i, b := range resp.Books {
		books[i] = entities.Book{
			ID:    int(b.Id),
			Title: b.Title,
			Author: entities.AuthorTable{
				ID:   int(b.Author.Id),
				Name: b.Author.FullName,
			},
		}
	}

	return books, nil
}

func (a *BookClient) GetBookWithAutor(id int) (entities.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var result entities.Book
	// Вызываем метод GetAllBookWithAuthor в сервисе Books
	resp, err := a.BooksServiceClient.GetBookWithAuthor(ctx, &proto.AuthorID{Id: int64(id)})
	if err != nil {
		return result, err
	}

	result.ID = int(resp.Id)
	result.Title = resp.Title
	result.Author = entities.AuthorTable{
		ID:   int(resp.Author.Id),
		Name: resp.Author.FullName,
	}
	return result, nil
}
func (a *BookClient) GetAllAuthorWithBooks() ([]entities.Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Вызываем метод GetAllAuthorWithBooks в сервисе Books
	resp, err := a.BooksServiceClient.GetAllAuthorWithBooks(ctx, &proto.Empty{})
	if err != nil {
		return nil, err
	}

	// Создаем слайс для хранения авторов
	authors := make([]entities.Author, len(resp.Authors))

	for i, a := range resp.Authors {
		books := make([]entities.BookTable, len(a.Books))
		for j, b := range a.Books {
			books[j] = entities.BookTable{
				ID:       int(b.Id),
				Title:    b.Title,
				AuthorID: int(b.AuthorId),
			}
		}

		authors[i] = entities.Author{
			ID:    int(a.Id),
			Name:  a.FullName,
			Books: books,
		}
	}

	return authors, nil
}
