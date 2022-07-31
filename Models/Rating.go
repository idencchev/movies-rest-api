package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingSchema struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	MovieId string             `json:"movieId,omitempty"`
	UserId  string             `json:"userId,omitempty"`
	Rating  int64              `json:"rating,omitempty"`
}
