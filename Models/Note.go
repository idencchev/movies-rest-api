package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoteSchema struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	MovieId string             `json:"movieId" bson:"movieId"`
	User    User               `json:"user" bson:"user"`
	Note    string             `json:"note" bson:"note"`
}

type User struct {
	Username string             `json:"username" bson:"username"`
	UserId   primitive.ObjectID `json:"userId" bson:"userId"`
}
