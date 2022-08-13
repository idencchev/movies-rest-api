package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"movies-rest-api/Models"
	"movies-rest-api/utils"

	"go.mongodb.org/mongo-driver/bson"
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
		_, err := collection.InsertOne(ctx, bson.M{"movieId": isRated.MovieId, "userId": isRated.UserId, "rating": isRated.Rating})
		if err != nil {
			w.Write([]byte(`{"message":"` + err.Error() + ` and"}`))
		}
		return
	}

	update := bson.M{
		"$set": bson.M{"movieId": isRated.MovieId, "userId": isRated.UserId, "rating": isRated.Rating},
	}

	err = collection.FindOneAndUpdate(ctx, filter, update).Decode(&dbIsRated)
	if err != nil {
		w.Write([]byte(`{"message":"` + err.Error() + ` and"}`))
	}
	return
}

func GetMovieRating(w http.ResponseWriter, r *http.Request) {

}
