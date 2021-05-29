package postgres

import (
	"testing"
)

func TestNewPgStoreFromConnString(t *testing.T) {
	_, err := NewPgStoreFromConnString("postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	if err != nil {
		t.Error(err)
	}
}

func TestPgStore_RunMigrations(t *testing.T) {
	store, _ := NewPgStoreFromConnString("postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	err := store.RunMigrations()
	if err != nil {
		t.Error(err)
	}
}
