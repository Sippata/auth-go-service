package network

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Sippata/auth-go-service/app"
)

// LoginHandler handles authentication requests
type LoginHandler struct {
	TokenService app.RefreshTokenService
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["client_id"]
	if !ok || len(keys[0]) < 1 {
		w.Write([]byte("Missing client id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userid := keys[0]

	tokenPair, err := app.CreateTokens(userid)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = h.TokenService.Add(tokenPair["refresh_token"], userid); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(tokenPair)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
