package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingSchema struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	MovieId string             `json:"movieId" bson:"movieId"`
	UserId  primitive.ObjectID `json:"userId" bson:"userId"`
	Rating  int64              `json:"rating" bson:"rating"`
}
