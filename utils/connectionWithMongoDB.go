package utils

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectionWithMongoDB(dbCollection string) (*mongo.Collection, context.Context, error) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://idenchev:ezhpclHpeWokofrR@cluster0.zusmc.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		// defer cancel()
		client, err := mongo.Connect(ctx, clientOptions)
		collection := client.Database("movies").Collection(dbCollection)
	if err != nil {
		log.Fatal(err)
	}
	return collection, ctx, err
}
