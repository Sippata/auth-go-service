package controllers

import (
	"encoding/json"
	"log"
	"medods-test/app"
	"net/http"
)

// Login a user in app
func Login(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["userid"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
	}
	userid := keys[0]

	td, err := app.CreateTokens(userid)
	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	response, err := json.Marshal(td)
	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(response)
}
