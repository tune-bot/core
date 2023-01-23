package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB = nil
var db_user = os.Getenv("DB_USER")
var db_pass = os.Getenv("DB_PASS")
var db_host = os.Getenv("DB_HOST")

func Connect() error {
	var err error
	db, err = sql.Open("mysql", db_user+":"+db_pass+"@tcp("+db_host+")/tune_bot")
	return err
}

func Disconnect() {
	if db != nil {
		db.Close()
		db = nil
	}
}
