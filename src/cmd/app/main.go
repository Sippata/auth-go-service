package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Sippata/medods-test/src/mongo"
	"github.com/Sippata/medods-test/src/network"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading: can't load env file")
	}

	client, err := mongo.Open()
	defer mongo.Close(client)
	dbName := os.Getenv("DB_NAME")

	tokenService := mongo.TokenService{DB: client.Database(dbName)}
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
	http.Handle("/refresh", &refreshHandler)
	http.Handle("/logout", &logoutHandler)
	http.Handle("/logout/all", &allLogoutHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
