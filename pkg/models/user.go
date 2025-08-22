package models

import (
	"context"
	"log"

	"github.com/Sambit99/Basic-RestAPI-Go-MongoDB/pkg/config"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var db *mongo.Collection
var ctx = context.Background()

type User struct {
	ID     bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string        `json:"name" bson:"name"`
	Age    int           `json:"age" bson:"age"`
	Gender string        `json:"gender" bson:"gender"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetAllUser() []bson.M {
	cursor, err := db.Find(ctx, bson.M{}, nil)

	if err != nil {
		log.Fatal("Error while getting all users", err.Error())
	}
	defer cursor.Close(ctx)

	var users []bson.M

	for cursor.Next(ctx) {
		var user bson.M

		if err = cursor.Decode(&user); err != nil {
			log.Fatal("Error while decoding user", err.Error())
		}

		users = append(users, user)
	}

	return users
}
