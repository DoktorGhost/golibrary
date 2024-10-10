package handlers

import (
	"github.com/DoktorGhost/golibrary/internal/providers"
	"github.com/DoktorGhost/golibrary/pkg/logger"
	"github.com/go-chi/chi"
)

func SetupRoutes(provider *providers.UseCaseProvider, logger logger.Logger) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/author/add", handlerAddAuthor(provider, logger))
	r.Post("/books/add", handlerAddBook(provider, logger))
	r.Get("/books", handlerGetAllBooks(provider, logger))
	r.Get("/authors", handlerGetAllAuthors(provider, logger))

	r.Post("/user/add", handlerAddUser(provider, logger))
	r.Get("/user/{id}", handlerGetUser(provider, logger))

	r.Get("/top/{period}/{limit}", handlerGetTop(provider, logger))
	r.Get("/rental", handlerGetAllRentals(provider, logger))
	r.Post("/rental/add/{user_id}/{book_id}", handlerGiveBook(provider, logger))
	r.Post("/rental/back/{book_id}", handlerBackBook(provider, logger))

	return r
}
