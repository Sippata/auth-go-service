package main

import (
	"fmt"
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

	var db mongo.MongoDB
	if err := db.Open(); err != nil {
		panic(fmt.Sprintf("Db connection fault: %v", err))
	}
	defer db.Close()

	tokenService := mongo.TokenService{
		DB:     db.Client.Database(os.Getenv("DB_NAME")),
		Client: db.Client,
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
	http.Handle("/refresh", &refreshHandler)
	http.Handle("/logout", &logoutHandler)
	http.Handle("/logout/all", &allLogoutHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
