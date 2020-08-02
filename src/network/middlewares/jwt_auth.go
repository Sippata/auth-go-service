package network

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Sippata/medods-test/src/app"
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
		}

		token, err := app.ParseToken(authHeader[1], []byte(os.Getenv("ACCESS_SECRET")))
		if claims, ok := token.Claims.(jwt.StandardClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "claims", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			log.Fatal(err)
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}
