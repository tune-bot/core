package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB = nil

func Connect(connection string) error {
	var err error
	db, err = sql.Open("mysql", connection)
	return err
}

func Disconnect() {
	if db != nil {
		db.Close()
		db = nil
	}
}
