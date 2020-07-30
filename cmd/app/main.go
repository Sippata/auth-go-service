package main

import (
	"log"
	"medods-test/network"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading: can't load env file")
	}

	var h network.LoginHandler
	h.TokenService = nil

	http.Handle("/login", &h)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
