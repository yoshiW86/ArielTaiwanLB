package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var SQLDB *sql.DB

func mysql() {
	var err error
	log.Println("DatabaseAC:", os.Getenv("DatabaseAC"), "DatabasePW:", os.Getenv("DatabasePW"), "dbURL:",os.Getenv("dbURL"), "database:",os.Getenv("database"))

	SQLDB, err = sql.Open("mysql", os.Getenv("DatabaseAC")+":"+os.Getenv("DatabasePW")+"@tcp("+os.Getenv("dbURL")+":3306)/"+os.Getenv("database"))

	if err != nil {
		log.Fatal(err.Error())
	}
	err = SQLDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
