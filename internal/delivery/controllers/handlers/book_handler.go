package handlers

import (
	"encoding/json"
	_ "github.com/DoktorGhost/golibrary/docs"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/entities"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/repositories/postgres/dao"
	"github.com/DoktorGhost/golibrary/internal/providers"
	"github.com/DoktorGhost/golibrary/pkg/logger"
	"io"
	"net/http"
	"strconv"
)

// @Summary Добавить автора
// @Description Добавляет нового автора в систему.
// @Tags Library
// @Accept json
// @Produce json
// @Param author body entities.AuthorRequest true "ФИО Автора"
// @Success 201 {string} string "Автор успешно добавлен, ID: {id}"
// @Failure 400 {string} string "Ошибка декодирования JSON или чтения тела запроса"
// @Failure 500 {string} string "Ошибка при добавлении автора"
// @Router /author/add [post]
// @Security BearerAuth
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

// @Summary Добавить книгу
// @Description Добавляет новую книгу в систему.
// @Tags Library
// @Accept json
// @Produce json
// @Param book body entities.BookRequest true "Информация о книге: название и ID автора"
// @Success 201 {string} string "Книга успешно добавлена, ID: {id}"
// @Failure 400 {string} string "Ошибка декодирования JSON или чтения тела запроса"
// @Failure 500 {string} string "Ошибка при добавлении книги"
// @Router /books/add [post]
// @Security BearerAuth
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

// @Summary Получить все книги
// @Description Возвращает список всех книг с информацией об авторах.
// @Tags Library
// @Accept json
// @Produce json
// @Success 200 {array} entities.Book "Список книг"
// @Failure 500 {string} string "Ошибка получения книг"
// @Router /books [get]
// @Security BearerAuth
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

// @Summary Получить всех авторов
// @Description Возвращает список всех авторов с информацией о их книгах.
// @Tags Library
// @Accept json
// @Produce json
// @Success 200 {array} entities.Author "Список авторов"
// @Failure 500 {string} string "Ошибка получения авторов"
// @Router /authors [get]
// @Security BearerAuth
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
