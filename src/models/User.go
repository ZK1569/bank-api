package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Firstname string             `bson:"firstname"`
	Age       int                `bson:"age"`
	History   []History          `bson:"history"`
}

type History struct {
	Money int `bson:"money"`
}
