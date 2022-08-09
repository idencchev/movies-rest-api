package utils

import (
	"github.com/go-chi/jwtauth"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var SECRET_KEY = []byte("gosecretkey")

var tokenAuth *jwtauth.JWTAuth

func GenerateJWT(userId primitive.ObjectID, username string) (string, error) {
	tokenAuth = jwtauth.New("HS256", []byte(SECRET_KEY), nil)
	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{"user_id": userId, "username": username})
	return tokenString, err
}
