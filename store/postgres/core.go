package postgres

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

type PgStore struct {
	db     *sql.DB
	driver database.Driver
}

func NewPgStoreFromConnString(connString string) (*PgStore, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	return NewPgStoreFromDB(db)
}

func NewPgStoreFromDB(db *sql.DB) (*PgStore, error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	return &PgStore{db: db, driver: driver}, nil
}

func (s *PgStore) RunMigrations() error {
	m, err := migrate.NewWithDatabaseInstance(
		"github://davevurby:davevurby@davevurby/mama-keeper/persistence/postgres/migrations",
		"postgres", s.driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		return err
	}

	return nil
}

func (s *PgStore) Close() error {
	return s.db.Close()
}
