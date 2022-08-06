package services

import (
	"encoding/json"
	"log"
	"net/http"
	
	"movies-rest-api/Models"
	"movies-rest-api/middlewares"
	"movies-rest-api/utils"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func UserSignup(w http.ResponseWriter, r *http.Request) {
	var client, ctx, _ = middlewares.ConnectionWithMongoDB()
	
	w.Header().Set("Content-Type", "application/json")

	var user models.UserSchema

	json.NewDecoder(r.Body).Decode(&user)

	user.Password = utils.GetHash([]byte(user.Password))
	collection := client.Database("movies").Collection("users")
	result, _ := collection.InsertOne(ctx, user)

	json.NewEncoder(w).Encode(result)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var client, ctx, err = middlewares.ConnectionWithMongoDB()
	
	w.Header().Set("Content-Type", "application/json")

	var user models.UserSchema
	var dbUser models.UserSchema

	json.NewDecoder(r.Body).Decode(&user)

	collection := client.Database("movies").Collection("users")
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
		w.Write([]byte(`{"response":"Wrong Password!"}`))
		return
	}

	jwtToken, err := utils.GenerateJWT()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	w.Write([]byte(`{"token":"` + jwtToken + `"}`))
}