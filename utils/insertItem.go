package utils

import (
	"context"
	"fmt"
	"log"
	"movies-rest-api/Models"

	"go.mongodb.org/mongo-driver/mongo"
)

// Collection object/instance
var collection *mongo.Collection

func InsertItem(user models.UserSchema) {
	insertResult, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}
