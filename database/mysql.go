package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func mysql() {
	var err error
	log.Println("DatabaseAC:", os.Getenv("DatabaseAC"), "DatabasePW:", os.Getenv("DatabasePW"), "dbURL:",os.Getenv("dbURL"), "database:",os.Getenv("database"))
	dbPath := os.Getenv("DatabaseAC")+":"+os.Getenv("DatabasePW")+"@tcp("+os.Getenv("dbURL")+":3306)/"+os.Getenv("database")
	SqlDB, err = sql.Open("mysql", dbPath)

	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
