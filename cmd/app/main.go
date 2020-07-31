package main

import (
	"log"
	"medods-test/mongo"
	"medods-test/network"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading: can't load env file")
	}

	client, err := mongo.Open()
	defer mongo.Close(client)
	dbName := os.Getenv("DB_NAME")

	var h network.LoginHandler
	h.TokenService = &mongo.TokenService{DB: client.Database(dbName)}

	http.Handle("/login", &h)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
