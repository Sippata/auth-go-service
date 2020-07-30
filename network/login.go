package network

import (
	"encoding/json"
	"log"
	"medods-test/app"
	"net/http"
)

// LoginHandler handles authentication requests
type LoginHandler struct {
	TokenService *app.TokenService
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["userid"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
	}
	userid := keys[0]

	tokensInfo, err := app.CreateTokens(userid)
	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	var tokenPair map[string]string
	tokenPair["access_token"] = tokensInfo["access_token"]
	tokenPair["refresh_token"] = tokensInfo["refresh_token"]
	response, err := json.Marshal(tokenPair)
	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(response)
}
