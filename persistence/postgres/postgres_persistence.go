package postgres

import (
	"database/sql"
	mama_keeper "github.com/davevurby/mama-keeper"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/lib/pq"
	"log"
)

type MKPostgresPersistence struct {
	db     *sql.DB
	driver database.Driver
}

func CreateMKPostgresPersistence(connString string) (*MKPostgresPersistence, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	mkp := &MKPostgresPersistence{db: db, driver: driver}
	err = mkp.runMigrations()
	if err != nil {
		log.Fatal(err)
	}

	return mkp, nil
}

func (mkp *MKPostgresPersistence) runMigrations() error {
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", mkp.driver)
	if err != nil {
		log.Fatal(err)
	}

	return m.Steps(1)
}

func (mkp *MKPostgresPersistence) CreateIdentity(identity mama_keeper.Identity) error {
	log.Printf("Upserting identity %s...\n", identity.Id)

	_, err := mkp.db.Exec("insert into identity (id, namespace, roles) values ($1, $2, $3) on conflict (id, namespace) do update set roles = $3", identity.Id, identity.Namespace, pq.Array(identity.Roles))
	return err
}
