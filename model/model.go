package model

import (
	"github.com/jinzhu/gorm"
	// gorm postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// Connect to the database or panic.
// If successful returns a database object.
func Connect(database string) *gorm.DB {
	var err error
	db, err = gorm.Open("postgres", "host=localhost user=postgres dbname="+database+" sslmode=disable password=admin port=5433")

	if err != nil {
		panic(err)
	}

	return db
}
