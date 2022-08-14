package services

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"movies-rest-api/Models"
	"movies-rest-api/utils"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func UserSignup(w http.ResponseWriter, r *http.Request) {
	var collection, ctx, _ = utils.ConnectionWithMongoDB("users")

	w.Header().Set("Content-Type", "application/json")

	var user models.UserSchema
	json.NewDecoder(r.Body).Decode(&user)

	user.Password = utils.GetHash([]byte(user.Password))
	result, _ := collection.InsertOne(ctx, bson.M{"username": user.Username, "password": user.Password})

	json.NewEncoder(w).Encode(result)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var collection, ctx, err = utils.ConnectionWithMongoDB("users")

	w.Header().Set("Content-Type", "application/json")

	var user models.UserSchema
	var dbUser models.UserSchema

	json.NewDecoder(r.Body).Decode(&user)
	err = collection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&dbUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)
	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

	if passErr != nil {
		log.Println(passErr)
		w.Write([]byte(`{"message":"Wrong Password!"}`))
		return
	}

	jwtToken, err := utils.GenerateJWT(dbUser.ID, dbUser.Username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	cookie := http.Cookie{Name: "x-auth-token", Value: jwtToken, Path: "/", Expires: time.Now().Add(2 * 24 * time.Hour), HttpOnly: true}
	http.SetCookie(w, &cookie)

	w.Write([]byte(`{"token":"` + jwtToken + `"}`))
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "x-auth-token",
		Path:   "/",
		MaxAge: -1,
	})
}

func Verify(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("x-auth-token")

	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			json.NewEncoder(w).Encode(false)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// For any other type of error, return a bad request status
		json.NewEncoder(w).Encode(false)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(true)
	return
}
