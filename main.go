package main

import (
	"github.com/IamNator/GoNepaApi/app"
	"github.com/IamNator/GoNepaApi/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func setupRouter(router mux.Router) {
	router.
		Methods("POST").
		Path("/endpoint").
		HandlerFunc(postFunction)
}

func main() {

	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	}

	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: database,
	}

	app.SetupRouter()

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
