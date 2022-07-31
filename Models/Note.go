package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoteSchema struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User   User               `json:"username,omitempty"`
	UserId string             `json:"userId,omitempty"`
	Rating int64              `json:"rating,omitempty"`
}

type User struct {
	Username string `json:"username,omitempty"`
	UserId string  `json:"userId,omitempty"`
}
