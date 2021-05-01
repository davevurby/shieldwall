package postgres

import (
	"database/sql"
)

type MKPostgresPersistence struct {
	db *sql.DB
}

func CreateMKPostgresPersistence(connString string) (*MKPostgresPersistence, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	return &MKPostgresPersistence{db: db}, nil
}
