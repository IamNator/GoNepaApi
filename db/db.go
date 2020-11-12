package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func CreateDatabase() (*sql.DB, error) {
	ServerName := "localhost:3306"
	user := "myuser"
	password := "pw"
	dbName := "nepa"

	connectingString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", user, password, ServerName, dbName)

	db, err := sql.Open("mysql", connectingString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
