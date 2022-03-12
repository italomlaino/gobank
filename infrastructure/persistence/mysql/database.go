package mysql

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Open() {
	dbDriver := "mysql"
	db, err := sql.Open(dbDriver, os.Getenv("DATASOURCE_URL"))
	if err != nil {
		panic(err.Error())
	}

	DB = db
}

func Close() {
	DB.Close()
	DB = nil
}
