package models

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	ID     bson.ObjectID `json:"id,omitempty" bson:"_id"`
	Name   string        `json:"name"`
	Age    int           `json:"age"`
	Gender string        `json:"gender"`
}
