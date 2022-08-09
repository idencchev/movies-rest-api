package middlewares

import (
	"context"
	"fmt"
	"movies-rest-api/services"
	"movies-rest-api/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenFromCookie, err := r.Cookie("x-auth-token")

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			services.Logout(w, r)
			w.Write([]byte("Malformed Token"))
		} else {

			jwtToken := tokenFromCookie.Value

			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(utils.SECRET_KEY), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "props", claims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				fmt.Println(err)
				services.Logout(w, r)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	})
}
