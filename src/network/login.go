package network

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Sippata/medods-test/src/app"
)

// LoginHandler handles authentication requests
type LoginHandler struct {
	TokenService app.RefreshTokenService
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["userid"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
	}
	userid := keys[0]

	tokenPair, err := app.CreateTokens(userid)
	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err = h.TokenService.Add(tokenPair["refresh_token"], userid); err != nil {
		log.Fatal(err)
	}

	response, err := json.Marshal(tokenPair)
	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(response)
}
