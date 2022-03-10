package mysql

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (db *sql.DB) {
	dbDriver := "mysql"
	db, err := sql.Open(dbDriver, os.Getenv("DATASOURCE_URL"))
	if err != nil {
		panic(err.Error())
	}
	return db
}
