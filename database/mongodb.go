package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client
var COLLECTION *mongo.Collection

func SetupMongo(mongoUri string) {

	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(mongoUri))
	if err != nil {
		panic(err)
	}

	coll := client.Database("golang").Collection("users")

	DB = client
	COLLECTION = coll
}
