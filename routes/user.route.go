package routes

import (
	"context"
	"encoding/json"
	"mongodb-server/database"
	"mongodb-server/types"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetUsersHandle(w http.ResponseWriter, r *http.Request) {

	var users types.User
	usersArray := []types.User{}

	cursor, err := database.COLLECTION.Find(context.TODO(), bson.D{})

	if err != nil {
		panic(err.Error())
	}

	for cursor.Next(context.TODO()) {
		if err := cursor.Decode(&users); err != nil {
			panic(err.Error())
		}
		usersArray = append(usersArray, users)

	}

	json.NewEncoder(w).Encode(&usersArray)

}
