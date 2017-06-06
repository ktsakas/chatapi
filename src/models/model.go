package model

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
)

// Connect to the database or panic.
// If successful returns a database object.
func Connect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=collegechat sslmode=disable password=admin port=5433")

	if err != nil {
		panic(err)
	}

	return db
}