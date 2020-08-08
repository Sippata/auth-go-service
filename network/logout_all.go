package network

import (
	"net/http"

	"github.com/Sippata/auth-go-service/app"

	"github.com/dgrijalva/jwt-go"
)

// AllLogout handle removing all refresh tokens for current user
type AllLogout struct {
	TokenService app.RefreshTokenService
}

func (h *AllLogout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	accessClaims := r.Context().Value("claims").(jwt.StandardClaims)

	err := h.TokenService.RemoveByUserID(accessClaims.Subject)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
