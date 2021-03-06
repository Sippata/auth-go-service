package network

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Sippata/auth-go-service/app"
	"github.com/dgrijalva/jwt-go"
)

// JwtMiddleware provides JSON Web Token authorization
func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			log.Print("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed token"))
			return
		}

		token, err := app.ParseToken(authHeader[1], []byte(os.Getenv("ACCESS_SECRET")))
		if err != nil {
			log.Print("Forbidden: " + err.Error())
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Token expired"))
			return
		}
		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			oldContext := r.Context()
			ctx := context.WithValue(oldContext, "claims", *claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			w.Write([]byte("Invalid token"))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	})
}
