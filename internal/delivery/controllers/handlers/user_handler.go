package handlers

import (
	"encoding/json"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_book/entities"
	entities2 "github.com/DoktorGhost/golibrary/internal/core/user/entities"
	"github.com/DoktorGhost/golibrary/internal/providers"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"io"
	"net/http"
	"strconv"
)

// @Summary Добавить пользователя
// @Description Добавляет нового пользователя в систему.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body entities.RegisterData true "Данные для регистрации пользователя"
// @Success 201 {string} string "Пользователь успешно добавлен, ID: {id}"
// @Failure 400 {string} string "Ошибка декодирования JSON или чтения тела запроса"
// @Failure 500 {string} string "Ошибка при добавлении пользователя"
// @Router /user/add [post]

func handlerAddUser(useCaseProvider *providers.UseCaseProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Неправильный метод", http.StatusMethodNotAllowed)
			return
		}

		// Чтение тела запроса
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
			log.Error("Ошибка чтения тела запроса", err)
			return
		}
		defer r.Body.Close()

		// Декодирование JSON из тела запроса
		var user entities.RegisterData
		if err := json.Unmarshal(body, &user); err != nil {
			http.Error(w, "Ошибка декодирования JSON: "+err.Error(), http.StatusBadRequest)
			log.Error("Ошибка декодирования JSON", err)
			return
		}

		// Вызов метода добавления автора из юзкейса

		id, err := useCaseProvider.UserUseCase.AddUser(user)
		if err != nil {
			http.Error(w, "Ошибка при добавлении пользователя: "+err.Error(), http.StatusInternalServerError)
			log.Error("Ошибка при добавлении пользователя", err)
			return
		}

		// Успешный ответ
		w.WriteHeader(http.StatusCreated)
		responseMessage := "Пользователь успешно добавлен, ID: " + strconv.Itoa(id)
		w.Write([]byte(responseMessage))

		log.Info("Пользователь успешно добавлен", "id", id)

	}
}

// @Summary Получить пользователя
// @Description Возвращает информацию о пользователе по его ID.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} dao.UserTable "Информация о пользователе"
// @Failure 400 {string} string "Неверный ID"
// @Failure 500 {string} string "Ошибка при получении пользователя"
// @Router /user/{id} [get]
// @Security BearerAuth
func handlerGetUser(useCaseProvider *providers.UseCaseProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Неправильный метод", http.StatusMethodNotAllowed)
			return
		}

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Неверный ID", http.StatusBadRequest)
			log.Error("Неверный ID", err)
			return
		}

		// Получаем пользователя с помощью юзкейса
		user, err := useCaseProvider.UserUseCase.GetUserByID(id)
		if err != nil {
			http.Error(w, "Ошибка при получении пользователя: "+err.Error(), http.StatusInternalServerError)
			log.Error("Ошибка при получении пользователя:", err)
			return
		}

		// Отправляем успешный ответ с пользователем
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
		log.Info("Пользователя успешно получен", "id", id)

	}
}

// @Summary Логин пользователя
// @Description Аутентификация пользователя по имени пользователя и паролю, возвращает JWT-токен.
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body entities2.Login true "Данные для входа"
// @Success 200 {object} map[string]string "JWT-токен"
// @Failure 400 {string} string "Ошибка декодирования данных или ошибка аутентификации"
// @Failure 405 {string} string "Неправильный метод"
// @Router /login [post]
func handlerLogin(useCaseProvider *providers.UseCaseProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Неправильный метод", http.StatusMethodNotAllowed)
			return
		}

		// Проверка на пустое тело запроса
		if r.Body == nil {
			http.Error(w, "Отсутствует тело запроса", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var loginData entities2.Login
		err := json.NewDecoder(r.Body).Decode(&loginData)
		if err != nil {
			http.Error(w, "Ошибка декодирования", http.StatusBadRequest)
			log.Error("Ошибка декодирования", err)
			return
		}

		var token string
		tokenString := r.Header.Get("Authorization")

		// Если токен есть в заголовке
		if tokenString != "" {
			// Проверка, начинается ли токен с "Bearer "
			if len(tokenString) > len("Bearer ") && tokenString[:len("Bearer ")] == "Bearer " {
				// Удаляем "Bearer " из строки токена
				tokenString = tokenString[len("Bearer "):]

				// Проверка токена
				tokenJWT, err := jwtauth.VerifyToken(useCaseProvider.AuthUseCase.TokenAuth, tokenString)
				if err != nil {
					log.Error("Неверный токен", err)
				} else {
					claims := tokenJWT.PrivateClaims()
					username, ok := claims["username"].(string)
					if ok && username == loginData.Username {
						// Токен действителен и соответствует пользователю
						token = tokenString
					}
				}
			} else {
				http.Error(w, "Неверный формат токена", http.StatusUnauthorized)
				return
			}
		}

		// Если токен не найден, аутентификация пользователя
		if token == "" {
			token, err = useCaseProvider.AuthUseCase.Login(loginData.Username, loginData.Password)
			if err != nil {
				http.Error(w, "Ошибка аутентификации", http.StatusBadRequest)
				log.Error("Ошибка аутентификации", err)
				return
			}
		} else {
			log.Error("Используем старый токен", err)
		}

		// Успешная аутентификация — возвращаем токен
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]string{"token": token}
		json.NewEncoder(w).Encode(response)
	}
}
