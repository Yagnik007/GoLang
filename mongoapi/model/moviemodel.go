package moviemodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	//Authogenerated id from database
	Movie   string `json:"movie,omitempty"`
	Watched bool   `json:"watched,omitempty"`
}
