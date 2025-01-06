package routes

import (
	"context"
	"encoding/json"
	"mongodb-server/database"
	"mongodb-server/types"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func FindUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(vars["id"])

	if err != nil {
		panic(err.Error())
	}

	findQuery := bson.M{"_id": id}
	var user types.User

	if err := database.COLLECTION.FindOne(context.TODO(), findQuery).Decode(&user); err != nil {
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}
