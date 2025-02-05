package db

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// open db connection and retuns
func Connect() (*sql.DB, error) {
	db, error := sql.Open("mysql", config.ConnectionString)
	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		db.Close()
		return nil, error
	}

	return db, nil

}
