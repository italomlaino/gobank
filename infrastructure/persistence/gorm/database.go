package gorm

import (
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Open() (db *gorm.DB, err error) {
	dsn := os.Getenv("DATASOURCE_URL")
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NameReplacer: strings.NewReplacer("Gorm", ""),
		  },
	})
}
