package migration

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

func Setup() {
	db, _ := sql.Open("postgres", "postgres://localhost:5432/database?sslmode=enable")
	driver, _ := postgres.WithInstance(db, &postgres.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"./setup.sql",
		"postgres", driver)

	m.Steps(1)
}
