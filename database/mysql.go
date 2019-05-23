package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func init() {
	var err error
	// dbPath := os.Getenv("DatabaseAC")+":"+os.Getenv("DatabasePW")+"@tcp("+os.Getenv("dbURL")+":3306)/"+os.Getenv("database")
	SqlDB, err = sql.Open("mysql", os.Getenv("DatabaseAC")+":"+os.Getenv("DatabasePW")+"@tcp("+os.Getenv("dbURL")+":3306)/"+os.Getenv("database"))

	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
