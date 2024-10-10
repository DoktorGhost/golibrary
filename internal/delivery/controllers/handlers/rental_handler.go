package handlers

import (
	"encoding/json"
	"github.com/DoktorGhost/golibrary/internal/providers"
	"github.com/DoktorGhost/golibrary/pkg/logger"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func handlerGetAllRentals(useCaseProvider *providers.UseCaseProvider, logger logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Неправильный метод", http.StatusMethodNotAllowed)
			return
		}

		rentals, err := useCaseProvider.LibraryUseCase.GetUserRentals()
		if err != nil {
			http.Error(w, "Ошибка чтения аренды: "+err.Error(), http.StatusInternalServerError)
			logger.Error("Ошибка чтения аренды", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(rentals); err != nil {
			http.Error(w, "Ошибка кодирования ответа: "+err.Error(), http.StatusInternalServerError)
			logger.Error("Ошибка кодирования ответа", err)
			return
		}

		logger.Info("Запрос на получение всех записей аренды успешно выполнен")

	}
}

func handlerGiveBook(useCaseProvider *providers.UseCaseProvider, logger logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Неправильный метод", http.StatusMethodNotAllowed)
			return
		}

		// Извлечение параметров из URL
		userIdStr := chi.URLParam(r, "user_id")
		bookIdStr := chi.URLParam(r, "book_id")

		userID, err := strconv.Atoi(userIdStr)
		if err != nil {
			http.Error(w, "Неправильный UserID", http.StatusMethodNotAllowed)
			return
		}

		bookID, err := strconv.Atoi(bookIdStr)
		if err != nil {
			http.Error(w, "Неправильный BookID", http.StatusMethodNotAllowed)
			return
		}

		// Здесь должна быть логика выдачи книги
		rentalID, err := useCaseProvider.LibraryUseCase.GiveBook(bookID, userID)
		if err != nil {
			http.Error(w, "Ошибка при выдаче книги: "+err.Error(), http.StatusInternalServerError)
			logger.Error("Ошибка при выдаче книги", err)
			return
		}

		// Успешный ответ
		w.WriteHeader(http.StatusOK)

		responseMessage := "Книга успешно выдана, RentalID: " + strconv.Itoa(rentalID)
		w.Write([]byte(responseMessage))
		logger.Info("Книга успешно выдана", "rentalID", rentalID, "userID", userID, "bookID", bookID)
	}
}

func handlerBackBook(useCaseProvider *providers.UseCaseProvider, logger logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Неправильный метод", http.StatusMethodNotAllowed)
			return
		}

		// Извлечение параметров из URL
		bookIdStr := chi.URLParam(r, "book_id")

		bookID, err := strconv.Atoi(bookIdStr)
		if err != nil {
			http.Error(w, "Неправильный BookID", http.StatusMethodNotAllowed)
			return
		}

		// Здесь должна быть логика выдачи книги
		err = useCaseProvider.LibraryUseCase.BackBook(bookID)
		if err != nil {
			http.Error(w, "Ошибка при возврате книги: "+err.Error(), http.StatusInternalServerError)
			logger.Error("Ошибка при возврате книги", err)
			return
		}

		// Успешный ответ
		w.WriteHeader(http.StatusOK)

		w.Write([]byte("Книга успешно возвращена"))
		logger.Info("Книга успешно возвращена", "bookID", bookID)
	}
}

func handlerGetTop(useCaseProvider *providers.UseCaseProvider, logger logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Неправильный метод", http.StatusMethodNotAllowed)
			return
		}

		// Извлечение параметров из URL
		periodStr := chi.URLParam(r, "period")
		limitStr := chi.URLParam(r, "limit")

		period, err := strconv.Atoi(periodStr)
		if err != nil {
			http.Error(w, "Неправильный период", http.StatusMethodNotAllowed)
			return
		}

		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "Неправильный лимит", http.StatusMethodNotAllowed)
			return
		}

		topAuthors, err := useCaseProvider.LibraryUseCase.GetTopAuthors(period, limit)
		if err != nil {
			http.Error(w, "Ошибка получения топа авторов: "+err.Error(), http.StatusInternalServerError)
			logger.Error("Ошибка получения топа авторов", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(topAuthors); err != nil {
			http.Error(w, "Ошибка кодирования ответа: "+err.Error(), http.StatusInternalServerError)
			logger.Error("Ошибка кодирования ответа", err)
			return
		}

		logger.Info("Запрос на получение топа авторов успешно выполнен")

	}
}
