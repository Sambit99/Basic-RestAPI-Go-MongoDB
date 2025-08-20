package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var db *mongo.Collection
var connectionString = "mongodb://root:root@localhost:27017/?authSource=admin"
var dbName = "test"
var colName = "user"

func Connect() {
	clientOptions := options.Client().ApplyURI(connectionString)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	fmt.Println("✅ Connected to MongoDB!")

	collection := client.Database(dbName).Collection(colName)
	db = collection
}

func GetDB() *mongo.Collection {
	return db
}
