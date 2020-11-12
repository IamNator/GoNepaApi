package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (app *App) SetupRouter() {
	app.Router.
		Methods("POST").
		Path("/endpoint").
		HandleFunc(app.postFunction)
}

func (app *App) getFunction

func (app *App) postFunction(res http.ResponseWriter, req *http.Request) {
	_, err := app.Database.Exec("INSERT INTO `test` (name) VALUES ('myname')")
	if err != nil {
		log.Fatal("Database INSERT failed")
	}

	log.Println("You called a thing")
	res.WriteHeader(http.StatusOK)
}
