package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Sippata/auth-go-service/mongo"
	"github.com/Sippata/auth-go-service/network"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		// log.Fatal("Error loading: can't load env file")
		panic(err)
	}

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

	log.Print("Run server on http://127.0.0.1:3000/")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
