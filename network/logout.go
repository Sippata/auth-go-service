package network

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Sippata/auth-go-service/app"

	"github.com/dgrijalva/jwt-go"
)

// LogoutHandler handle refresh token removeing
type LogoutHandler struct {
	TokenService app.RefreshTokenService
}

func (h *LogoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body requestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.Write([]byte("Missing refresh token"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rt, err := app.ParseToken(body.Token, []byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	claims, ok := rt.Claims.(*jwt.StandardClaims)
	if !rt.Valid || !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.TokenService.Remove(claims.Id)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
