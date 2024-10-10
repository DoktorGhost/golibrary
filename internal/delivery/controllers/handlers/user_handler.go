package handlers

import (
	"encoding/json"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/entities"
	"github.com/DoktorGhost/golibrary/internal/providers"
	"github.com/DoktorGhost/golibrary/pkg/logger"
	"github.com/go-chi/chi"
	"io"
	"net/http"
	"strconv"
)

func handlerAddUser(useCaseProvider *providers.UseCaseProvider, logger logger.Logger) http.HandlerFunc {
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
		var user entities.RegisterData
		if err := json.Unmarshal(body, &user); err != nil {
			http.Error(w, "Ошибка декодирования JSON: "+err.Error(), http.StatusBadRequest)
			logger.Error("Ошибка декодирования JSON", err)
			return
		}

		// Вызов метода добавления автора из юзкейса

		id, err := useCaseProvider.UserUseCase.AddUser(user)
		if err != nil {
			http.Error(w, "Ошибка при добавлении пользователя: "+err.Error(), http.StatusInternalServerError)
			logger.Error("Ошибка при добавлении пользователя", err)
			return
		}

		// Успешный ответ
		w.WriteHeader(http.StatusCreated)
		responseMessage := "Пользователь успешно добавлен, ID: " + strconv.Itoa(id)
		w.Write([]byte(responseMessage))

		logger.Info("Пользователь успешно добавлен", "id", id)

	}
}

func handlerGetUser(useCaseProvider *providers.UseCaseProvider, logger logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Неправильный метод", http.StatusMethodNotAllowed)
			return
		}

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Неверный ID", http.StatusBadRequest)
			logger.Error("Неверный ID", err)
			return
		}

		// Получаем пользователя с помощью юзкейса
		user, err := useCaseProvider.UserUseCase.GetUserByID(id)
		if err != nil {
			http.Error(w, "Ошибка при получении пользователя: "+err.Error(), http.StatusInternalServerError)
			logger.Error("Ошибка при получении пользователя:", err)
			return
		}

		// Отправляем успешный ответ с пользователем
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
		logger.Info("Пользователя успешно получен", "id", id)

	}
}
