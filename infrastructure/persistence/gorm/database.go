package gorm

import (
	"database/sql"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var SqlDB *sql.DB

func Open() {
	dsn := os.Getenv("DATASOURCE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NameReplacer: strings.NewReplacer("Gorm", ""),
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	DB = db
	SqlDB = sqlDB
}

func Close() {
	SqlDB.Close()
	DB = nil
	SqlDB = nil
}
