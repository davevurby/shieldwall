package postgres

import (
	"testing"
)

func TestNewPostgresPersistence(t *testing.T) {
	_, err := NewPostgresPersistence("postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	if err != nil {
		t.Error(err)
	}
}

func TestMKPostgresPersistence_RunMigrations(t *testing.T) {
	mkp, _ := NewPostgresPersistence("postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	err := mkp.RunMigrations()
	if err != nil {
		t.Error(err)
	}
}
