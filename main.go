package main

import (
	"log"
	"mongodb-server/database"
	"mongodb-server/routes"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Panic(err.Error())
	}

	MONGO_URI := os.Getenv("MONGODB_URI")

	database.SetupMongo(MONGO_URI)

	r := mux.NewRouter()

	r.HandleFunc("/user", routes.GetUsersHandle)

	http.ListenAndServe(":3000", r)
}
