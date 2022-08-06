package middlewares

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection object/instance
var Collection *mongo.Collection

func ConnectionWithMongoDB() (*mongo.Client, context.Context, error) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://idenchev:ezhpclHpeWokofrR@cluster0.zusmc.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client, ctx, err
}
