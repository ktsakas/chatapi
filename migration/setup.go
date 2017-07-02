package migration

import (
	"database/sql"
	"log"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"

	// Required to avoid unknown driver file error
	_ "github.com/mattes/migrate/source/file"
)

// TODO: change this to use and sql dump to load database
func Rebuild() {
	db, err := sql.Open("postgres", "postgres://localhost:5433/collegechat?sslmode=disable&user=postgres&password=admin")
	if err != nil {
		println("Failed to connect to test database " + "collegechat.")
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		println("Failed to get driver from database instance " + "collegechat.")
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../migration",
		"postgres", driver)
	defer m.Close()
	if err != nil {
		println("Failed to run migration.")
		log.Fatal(err)
	}

	m.Drop()
	m.Up()
}
