package network

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Sippata/auth-go-service/app"

	"github.com/dgrijalva/jwt-go"
)

// RefreshHandler handle request and refresh token
type RefreshHandler struct {
	TokenService app.RefreshTokenService
}

type requestBody struct {
	Token string `json:"refresh_token"`
}

func (h *RefreshHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	accessClaims := r.Context().Value("claims").(jwt.StandardClaims)

	var body requestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.Write([]byte("Missing refresh token"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check that the refresh token is valid
	rToken, err := app.ParseToken(body.Token, []byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		w.Write([]byte("Malformed refresh token"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rtClaims, ok := rToken.Claims.(*jwt.StandardClaims)
	if !ok || !rToken.Valid {
		w.Write([]byte("Invalid refresh token"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check that the refresh token associate with the access token
	tokenHash, err := h.TokenService.Get(rtClaims.Id)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = app.Compare(tokenHash, body.Token)
	if err != nil || rtClaims.Audience != accessClaims.Id {
		w.Write([]byte("Given refresh token does not associate with this access token"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.TokenService.Remove(rtClaims.Id)

	// Create new token pair
	tokenPair, err := app.CreateTokens(accessClaims.Subject)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	rt, _ := app.ParseToken(tokenPair["refresh_token"], []byte(os.Getenv("REFRESH_SECRET")))
	if err = h.TokenService.Add(rt); err != nil {
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
