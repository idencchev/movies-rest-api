package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserSchema struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
