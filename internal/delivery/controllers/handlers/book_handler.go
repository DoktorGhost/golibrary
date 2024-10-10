package handlers

import (
	"encoding/json"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/entities"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/repositories/postgres/dao"
	"github.com/DoktorGhost/golibrary/internal/providers"
	"github.com/DoktorGhost/golibrary/pkg/logger"
	"io"
	"net/http"
	"strconv"
)

func handlerAddAuthor(useCaseProvider *providers.UseCaseProvider, logger logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Неправильный метод", http.StatusMethodNotAllowed)
			return
		}

		// Чтение тела запроса
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
			logger.Error("Ошибка чтения тела запроса", err)
			return
		}
		defer r.Body.Close()

		// Декодирование JSON из тела запроса
		var author entities.AuthorRequest
		if err := json.Unmarshal(body, &author); err != nil {
			http.Error(w, "Ошибка декодирования JSON: "+err.Error(), http.StatusBadRequest)
			logger.Error("Ошибка декодирования JSON", err)
			return
		}

		id, err := useCaseProvider.BookUseCase.AddAuthor(author.Name, author.Surname, author.Patronymic)
		if err != nil {
			http.Error(w, "Ошибка при добавлении автора: "+err.Error(), http.StatusInternalServerError)
			logger.Error("Ошибка при добавлении автора", err)
			return
		}

		// Успешный ответ
		w.WriteHeader(http.StatusCreated)
		responseMessage := "Автор успешно добавлен, ID: " + strconv.Itoa(id)
		w.Write([]byte(responseMessage))

		logger.Info("Автор успешно добавлен", "id", id)

	}
}

func handlerAddBook(useCaseProvider *providers.UseCaseProvider, logger logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Неправильный метод", http.StatusMethodNotAllowed)
			return
		}

		// Чтение тела запроса
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
			logger.Error("Ошибка чтения тела запроса", err)
			return
		}
		defer r.Body.Close()

		// Декодирование JSON из тела запроса
		var book entities.BookRequest
		if err := json.Unmarshal(body, &book); err != nil {
			http.Error(w, "Ошибка декодирования JSON: "+err.Error(), http.StatusBadRequest)
			logger.Error("Ошибка декодирования JSON", err)
			return
		}

		id, err := useCaseProvider.BookUseCase.AddBook(dao.BookTable{Title: book.Title, AuthorID: book.AuthorID})
		if err != nil {
			http.Error(w, "Ошибка при добавлении книги: "+err.Error(), http.StatusInternalServerError)
			logger.Error("Ошибка при добавлении книги", err)
			return
		}

		// Успешный ответ
		w.WriteHeader(http.StatusCreated)
		responseMessage := "Книга успешно добавлена, ID: " + strconv.Itoa(id)
		w.Write([]byte(responseMessage))

		logger.Info("Книга успешно добавлена", "id", id)

	}
}

func handlerGetAllBooks(useCaseProvider *providers.UseCaseProvider, logger logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Неправильный метод", http.StatusMethodNotAllowed)
			return
		}

		books, err := useCaseProvider.BookUseCase.GetAllBookWithAuthor()
		if err != nil {
			http.Error(w, "Ошибка получения книг: "+err.Error(), http.StatusInternalServerError)
			logger.Error("Ошибка получения книг", err)
			return
		}

		// Установите заголовок Content-Type
		w.Header().Set("Content-Type", "application/json")

		// Преобразуйте книги в JSON
		response, err := json.Marshal(books)
		if err != nil {
			http.Error(w, "Ошибка при преобразовании данных в JSON: "+err.Error(), http.StatusInternalServerError)
			logger.Error("Ошибка при преобразовании данных в JSON", err)
			return
		}

		w.Write(response)
		logger.Info("Книги успешно получены", "count", len(books))
	}
}

func handlerGetAllAuthors(useCaseProvider *providers.UseCaseProvider, logger logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Неправильный метод", http.StatusMethodNotAllowed)
			return
		}

		books, err := useCaseProvider.BookUseCase.GetAllAuthorWithBooks()
		if err != nil {
			http.Error(w, "Ошибка получения авторов: "+err.Error(), http.StatusInternalServerError)
			logger.Error("Ошибка получения авторов", err)
			return
		}

		// Установите заголовок Content-Type
		w.Header().Set("Content-Type", "application/json")

		// Преобразуйте книги в JSON
		response, err := json.Marshal(books)
		if err != nil {
			http.Error(w, "Ошибка при преобразовании данных в JSON: "+err.Error(), http.StatusInternalServerError)
			logger.Error("Ошибка при преобразовании данных в JSON", err)
			return
		}

		w.Write(response)

		logger.Info("Авторы успешно получены", "count", len(books))
	}
}
