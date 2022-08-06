package utils

import (
	"log"

	"github.com/dgrijalva/jwt-go"
)

var SECRET_KEY = []byte("gosecretkey")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}
