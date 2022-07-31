package utils

import (
	"context"
	"fmt"
	"log"
	"movies-rest-api/Models"
	"movies-rest-api/middlewares"
)

func InsertItem(user models.UserSchema) {
	insertResult, err := middlewares.Collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}
