package network

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Sippata/auth-go-service/app"

	"github.com/dgrijalva/jwt-go"
)

// LogoutHandler handle refresh token removeing
type LogoutHandler struct {
	TokenService app.RefreshTokenService
}

func (h *LogoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	accessClaims := r.Context().Value("claims").(jwt.StandardClaims)

	var body requestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.Write([]byte("Missing refresh token"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.TokenService.Remove(body.Token, accessClaims.Subject)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
