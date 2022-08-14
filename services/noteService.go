package services

import (
	"encoding/json"
	"fmt"
	"movies-rest-api/Models"
	"movies-rest-api/utils"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var collection, ctx, err = utils.ConnectionWithMongoDB("notes")

	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}

	w.Header().Set("Content-Type", "application/json")

	var note models.NoteSchema
	json.NewDecoder(r.Body).Decode(&note)

	insert := bson.M{"userId": note.User.UserId, "movieId": note.MovieId, "username": note.User.Username, "note": note.Note}
	result, err := collection.InsertOne(ctx, insert)

	if err != nil {
		fmt.Println("InsertOne ERROR:", err)
		os.Exit(1)
	}

	note.ID = result.InsertedID.(primitive.ObjectID)
	json.NewEncoder(w).Encode(note)
}

func GetNoteByMovieId(w http.ResponseWriter, r *http.Request) {
	var collection, ctx, err = utils.ConnectionWithMongoDB("notes")
	
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}

	w.Header().Set("Content-Type", "application/json")
	
	movieId := chi.URLParam(r, "movieId")
	

	filter := bson.M{"movieId": movieId}

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

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	var collection, ctx, err = utils.ConnectionWithMongoDB("notes")

	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}

	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")
	noteId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	collection.FindOneAndDelete(ctx, bson.M{"_id": noteId})
}
