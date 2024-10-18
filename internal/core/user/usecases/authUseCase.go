package usecases

import (
	"errors"
	"github.com/DoktorGhost/golibrary/internal/auth"
	"github.com/go-chi/jwtauth"

	"github.com/DoktorGhost/golibrary/internal/core/user/services"
)

type AuthUseCase struct {
	userService *services.UserService
	TokenAuth   *jwtauth.JWTAuth
}

func NewAuthUseCase(userService *services.UserService, key string) *AuthUseCase {
	token := jwtauth.New("HS256", []byte(key), nil)
	return &AuthUseCase{userService: userService, TokenAuth: token}
}

func (auc *AuthUseCase) Login(username, password string) (string, error) {
	user, err := auc.userService.GetUserByUsername(username)
	if err != nil {
		if err != nil {
			return "", errors.New("user not found")
		}
	}

	flag, err := auth.CheckPasswordHash(password, user.PasswordHash)
	if err != nil {
		return "", err
	}

	if !flag {
		return "", errors.New("invalid password")
	}

	jwt, err := auth.GenerateJWT(username, auc.TokenAuth)

	// Авторизация успешна, возвращаем JWT
	return jwt, nil
}
