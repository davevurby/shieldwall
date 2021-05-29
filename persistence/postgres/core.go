package postgres

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

type MKPostgresPersistence struct {
	db     *sql.DB
	driver database.Driver
}

func NewPostgresPersistence(connString string) (*MKPostgresPersistence, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	return NewPostgresPersistenceWithDB(db)
}

func NewPostgresPersistenceWithDB(db *sql.DB) (*MKPostgresPersistence, error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	return &MKPostgresPersistence{db: db, driver: driver}, nil
}

func (mkp *MKPostgresPersistence) RunMigrations() error {
	m, err := migrate.NewWithDatabaseInstance(
		"github://davevurby:davevurby@davevurby/mama-keeper/persistence/postgres/migrations",
		"postgres", mkp.driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		return err
	}

	return nil
}

func (mkp *MKPostgresPersistence) Close() error {
	return mkp.db.Close()
}
