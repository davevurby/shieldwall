package postgres

import (
	"testing"

	mama_keeper "github.com/davevurby/mama-keeper"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestMKPostgresPersistence_CreateRole(t *testing.T) {
	mkp, _ := NewPostgresPersistence("postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	err := mkp.CreateRole(mama_keeper.Role{Id: "test_role", Namespaces: []string{"mamakeeper.io/users", "mamakeeper.io/admins"}})
	if err != nil {
		t.Error(err)
	}

	err = mkp.CreateRole(mama_keeper.Role{Id: "test_role", Namespaces: []string{"mamakeeper.io/users", "mamakeeper.io/admins", "mamakeeper.io/companies/*/users"}})
	if err != nil {
		t.Error(err)
	}

	rows, err := mkp.db.Query("select * from role where id = $1", "test_role")
	if err != nil {
		t.Error(err)
	}

	result := rows.Next()
	if result == false {
		t.Error("there is no role")
	}

	var id string
	var namespaces []string
	if err = rows.Scan(&id, pq.Array(&namespaces)); err != nil {
		t.Error(err)
	}

	result = rows.Next()
	if result == true {
		t.Error("there should be no more role")
	}

	assert.Equal(t, id, "test_role", "it should return id as 'role_id'")
	assert.Equal(t, namespaces, []string{"mamakeeper.io/users", "mamakeeper.io/admins", "mamakeeper.io/companies/*/users"}, "it should return namespaces as 'mamakeeper.io/users', 'mamakeeper.io/admins' and 'mamakeeper.io/companies/*/users'")
}
