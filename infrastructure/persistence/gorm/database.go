package gorm

import (
	"database/sql"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Open() (*gorm.DB, *sql.DB, error) {
	dsn := os.Getenv("DATASOURCE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NameReplacer: strings.NewReplacer("Gorm", ""),
		},
	})
	if err != nil {
		return nil, nil, err
	}
	sqlDB, err := db.DB()
	return db, sqlDB, err
}
