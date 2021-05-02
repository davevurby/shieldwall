package postgres

import (
	"testing"
)

func TestCreateMKPostgresPersistence(t *testing.T) {
	_, err := CreateMKPostgresPersistence("postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	if err != nil {
		t.Error(err)
	}
}

//func TestMigration(t *testing.T) {
//	mkp, _ := CreateMKPostgresPersistence("postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
//	err := mkp.RunMigrations()
//	if err != nil {
//		t.Error(err)
//	}
//}
