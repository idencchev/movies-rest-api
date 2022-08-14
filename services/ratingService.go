package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"movies-rest-api/Models"
	"movies-rest-api/utils"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddMovieRating(w http.ResponseWriter, r *http.Request) {
	var collection, ctx, err = utils.ConnectionWithMongoDB("ratings")

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")

	var isRated models.RatingSchema
	var dbIsRated models.RatingSchema

	json.NewDecoder(r.Body).Decode(&isRated)

	filter := bson.M{"movieId": isRated.MovieId, "userId": isRated.UserId}

	result, err := collection.Find(ctx, filter)

	var results []bson.M
	if err = result.All(ctx, &results); err != nil {
		w.Write([]byte(`{"message":"` + err.Error() + ` and"}`))
	}

	if len(results) == 0 {
		createdItem, err := collection.InsertOne(ctx, bson.M{"movieId": isRated.MovieId, "userId": isRated.UserId, "rating": isRated.Rating})
		if err != nil {
			w.Write([]byte(`{"message":"` + err.Error() + ` and"}`))
			return
		}
		json.NewEncoder(w).Encode(createdItem)
		return
	}

	update := bson.M{
		"$set": bson.M{"movieId": isRated.MovieId, "userId": isRated.UserId, "rating": isRated.Rating},
	}

	err = collection.FindOneAndUpdate(ctx, filter, update).Decode(&dbIsRated)
	if err != nil {
		w.Write([]byte(`{"message":"` + err.Error() + ` and"}`))
	}
	json.NewEncoder(w).Encode(dbIsRated)
}

func GetMovieRating(w http.ResponseWriter, r *http.Request) {
	var collection, ctx, err = utils.ConnectionWithMongoDB("ratings")

	movieId := chi.URLParam(r, "movieId")
	userIdString := chi.URLParam(r, "userId")
	
	userId, err := primitive.ObjectIDFromHex(userIdString)
	if err != nil {
		panic(err)
	}

	filter := bson.M{"movieId": movieId, "userId": userId}
	result, err := collection.Find(ctx, filter)
	if err != nil {
		w.Write([]byte(`{"message":"` + err.Error() + ` and"}`))
		json.NewEncoder(w).Encode("This item does not exist.")
		return
	}

	var results []bson.M
	if err = result.All(ctx, &results); err != nil {
		w.Write([]byte(`{"message":"` + err.Error() + ` and"}`))
	}
	json.NewEncoder(w).Encode(results)
}
