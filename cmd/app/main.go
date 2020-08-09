package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Sippata/auth-go-service/mongo"
	"github.com/Sippata/auth-go-service/network"
)

func main() {
	var instance mongo.DBInstance
	if err := instance.Open(); err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer instance.Close()

	dbName := os.Getenv("DB_NAME")

	tokenService := mongo.TokenService{
		Instance:   &instance,
		Collection: instance.GetCollection(dbName, "refrsh_tokens"),
	}
	loginHandler := network.LoginHandler{
		TokenService: &tokenService,
	}
	refreshHandler := network.RefreshHandler{
		TokenService: &tokenService,
	}
	logoutHandler := network.LogoutHandler{
		TokenService: &tokenService,
	}
	allLogoutHandler := network.AllLogout{
		TokenService: &tokenService,
	}

	http.Handle("/login", &loginHandler)
	http.Handle("/refresh", network.JwtMiddleware(&refreshHandler))
	http.Handle("/logout", network.JwtMiddleware(&logoutHandler))
	http.Handle("/logout/all", network.JwtMiddleware(&allLogoutHandler))

	port := os.Getenv("PORT")
	log.Print("Run server on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
