package usecases

import (
	"errors"
	"github.com/DoktorGhost/golibrary/internal/core/user/entities"
	"github.com/go-chi/jwtauth"
	"github.com/golang-jwt/jwt/v4"
	"time"

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

func (auc *AuthUseCase) Login(userData entities.Login) (string, error) {
	user, err := auc.userService.Login(userData)
	if err != nil {
		return "", errors.New("login filed " + err.Error())
	}

	jwt, err := generateJWT(user.Username, auc.TokenAuth)

	// Авторизация успешна, возвращаем JWT
	return jwt, nil
}

func generateJWT(username string, tokenAuth *jwtauth.JWTAuth) (string, error) {
	_, tokenString, err := tokenAuth.Encode(jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 1440).Unix(),
	})
	return tokenString, err
}
