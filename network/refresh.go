package network

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Sippata/auth-go-service/app"

	"github.com/dgrijalva/jwt-go"
)

// RefreshHandler handle request and refresh token
type RefreshHandler struct {
	TokenService app.RefreshTokenService
}

type requestBody struct {
	Token string `json:"refrsh_token"`
}

func (h *RefreshHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	accessClaims := r.Context().Value("claims").(jwt.StandardClaims)

	var body requestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	token, err := h.TokenService.Get(body.Token, accessClaims.Subject)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Given refresh token does not associate with this access token"))
	}

	h.TokenService.Remove(token, accessClaims.Subject)

	tokenPair, err := app.CreateTokens(accessClaims.Subject)
	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err = h.TokenService.Add(tokenPair["refresh_token"], accessClaims.Subject); err != nil {
		log.Fatal(err)
	}

	response, err := json.Marshal(tokenPair)
	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(response)
}
