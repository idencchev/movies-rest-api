package services

import (
	"encoding/json"

	"movies-rest-api/Models"
	"movies-rest-api/middlewares"
	"movies-rest-api/utils"
	"net/http"
)

func init() {
	middlewares.ConnectionWithMongoDB("users")
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.UserSchema
	_ = json.NewDecoder(r.Body).Decode(&user)

	utils.InsertItem(user)
	json.NewEncoder(w).Encode(user)
}
