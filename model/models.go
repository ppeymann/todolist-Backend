package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todos struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Titles string `json:"title,omitempty"`
	Todo string `json:"todo,omitempty"`
	Finished bool `json:"finished,omitempty"`
}