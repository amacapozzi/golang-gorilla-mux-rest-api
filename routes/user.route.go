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
	"go.mongodb.org/mongo-driver/mongo"
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
		response := types.ApiResponse{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	findQuery := bson.M{"_id": id}
	var user types.User

	if err := database.COLLECTION.FindOne(context.TODO(), findQuery).Decode(&user); err != nil {
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	var user types.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var existingUser types.User
	err := database.COLLECTION.FindOne(context.TODO(), bson.M{"username": user.Name}).Decode(&existingUser)
	if err == nil {
		response := types.ApiResponse{
			Message: "User already exists",
		}
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(&response)
		return
	}

	if err != mongo.ErrNoDocuments {
		http.Error(w, "Error checking user", http.StatusInternalServerError)
		return
	}

	_, err = database.COLLECTION.InsertOne(context.TODO(), bson.M{"username": user.Name})
	if err != nil {
		response := types.ApiResponse{
			Message: "Failed to create user: " + err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response)
		return
	}
	response := types.ApiResponse{
		Message: "User created successfully",
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&response)
}
