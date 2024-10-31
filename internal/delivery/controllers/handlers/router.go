package handlers

import (
	"github.com/DoktorGhost/golibrary/internal/metrics"
	"github.com/DoktorGhost/golibrary/internal/providers"
	"github.com/DoktorGhost/platform/logger"
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

		//запросы к внешним API
		r.Post("/author/add", handlerAddAuthor(provider))
		r.Post("/books/add", handlerAddBook(provider))
		r.Get("/books", handlerGetAllBooks(provider))
		r.Get("/authors", handlerGetAllAuthors(provider))
		r.Get("/user/{id}", handlerGetUser(provider))

		//запросы к БД
		r.Get("/rentals", handlerGetAllRentals(provider))
		r.Post("/rental/add/{user_id}/{book_id}", handlerGiveBook(provider))
		r.Post("/rental/back/{book_id}", handlerBackBook(provider))
	})

	//запросы к внешним API
	r.Post("/login", handlerLogin(provider))
	r.Post("/register", handlerAddUser(provider))

	//метрики
	r.Handle("/metrics", promhttp.Handler())
	r.Get("/debug/pprof/", PprofHandler)

	// Swagger
	r.Get("/swagger*", httpSwagger.WrapHandler)

	return r
}
