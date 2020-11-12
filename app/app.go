package app

import (
	"database/sql"
	"log"
	"net/http"

	"encoding/json"
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

func (app *App) getFunction(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		log.Fatal("No ID in the path")
	}

	dbdata := &DbData{}
	err := app.Database.QueryRow("SELECT id, data, name FROM `test` WHERE id = ?", id).Scan(&dbdata.ID, &dbdata.Date, &dbdata.Name)

	if err != nil {
		log.Fatal("Database SELECT failed")
	}

	log.Println("You fetched a thing!")
	res.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(res).Encode(dbdata); err != nil {
		panic(err)
	}

}

func (app *App) postFunction(res http.ResponseWriter, req *http.Request) {
	_, err := app.Database.Exec("INSERT INTO `test` (name) VALUES ('myname')")
	if err != nil {
		log.Fatal("Database INSERT failed")
	}

	log.Println("You called a thing")
	res.WriteHeader(http.StatusOK)
}
