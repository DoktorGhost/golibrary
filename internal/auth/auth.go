package auth

import (
	"github.com/go-chi/jwtauth"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GenerateJWT(username string, tokenAuth *jwtauth.JWTAuth) (string, error) {
	_, tokenString, err := tokenAuth.Encode(jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 1440).Unix(),
	})
	return tokenString, err
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}
