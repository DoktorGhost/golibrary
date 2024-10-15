package handlers

import (
	"github.com/DoktorGhost/golibrary/internal/metrics"
	"github.com/DoktorGhost/golibrary/internal/providers"
	"github.com/DoktorGhost/golibrary/pkg/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
)

var (
	log *logger.ZapLogger
)

func SetupRoutes(provider *providers.UseCaseProvider) *chi.Mux {
	r := chi.NewRouter()
	log, _ = logger.GetLogger()

	r.Use(metrics.RequestMetricsMiddleware)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(provider.AuthUseCase.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/author/add", handlerAddAuthor(provider))
		r.Post("/books/add", handlerAddBook(provider))
		r.Get("/books", handlerGetAllBooks(provider))
		r.Get("/authors", handlerGetAllAuthors(provider))

		r.Get("/rentals", handlerGetAllRentals(provider))
		r.Get("/top/{period}/{limit}", handlerGetTop(provider))
		r.Post("/rental/add/{user_id}/{book_id}", handlerGiveBook(provider))
		r.Post("/rental/back/{book_id}", handlerBackBook(provider))

		r.Get("/user/{id}", handlerGetUser(provider))

	})

	r.Post("/login", handlerLogin(provider))
	r.Post("/user/add", handlerAddUser(provider))

	//метрики
	r.Handle("/metrics", promhttp.Handler())
	r.Get("/debug/pprof/", PprofHandler)

	// Настройка Swagger
	r.Get("/swagger*", httpSwagger.WrapHandler)

	return r
}
