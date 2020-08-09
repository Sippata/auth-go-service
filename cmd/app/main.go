package main

import (
	"html/template"
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

	//apiDoc := http.FileServer(http.Dir("../../static"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		apiDoc, err := template.ParseFiles("static/api_docs.html")
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		apiDoc.Execute(w, nil)
	})
	http.Handle("/login", &loginHandler)
	http.Handle("/refresh", network.JwtMiddleware(&refreshHandler))
	http.Handle("/logout", network.JwtMiddleware(&logoutHandler))
	http.Handle("/logout/all", network.JwtMiddleware(&allLogoutHandler))

	port := os.Getenv("PORT")
	log.Print("Listening server on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
