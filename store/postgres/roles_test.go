package postgres

import (
	"testing"

	"github.com/davevurby/shieldwall"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestPgStore_GetRole(t *testing.T) {
	store, _ := NewPgStoreFromConnString("postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	err := store.PutRole(shieldwall.Role{Id: "test_role", Namespaces: []string{"shieldwall.io/users", "shieldwall.io/admins", "shieldwall.io/companies/*/users"}})
	if err != nil {
		t.Error(err)
	}

	role, err := store.GetRole("test_role")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, role.Id, "test_role", "it should return id as 'test_role'")
	assert.Equal(t, role.Namespaces, []string{"shieldwall.io/users", "shieldwall.io/admins", "shieldwall.io/companies/*/users"}, "it should return namespaces as 'shieldwall.io/users', 'shieldwall.io/admins' and 'shieldwall.io/companies/*/users'")
}

func TestPgStore_GetRoles(t *testing.T) {
	store, _ := NewPgStoreFromConnString("postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	err := store.PutRole(shieldwall.Role{Id: "test_role", Namespaces: []string{"shieldwall.io/users", "shieldwall.io/admins", "shieldwall.io/companies/*/users"}})
	if err != nil {
		t.Error(err)
	}

	roles, err := store.GetRoles()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, roles[0].Id, "test_role", "it should return id as 'test_role'")
	assert.Equal(t, roles[0].Namespaces, []string{"shieldwall.io/users", "shieldwall.io/admins", "shieldwall.io/companies/*/users"}, "it should return namespaces as 'shieldwall.io/users', 'shieldwall.io/admins' and 'shieldwall.io/companies/*/users'")
}

func TestPgStore_PutRole(t *testing.T) {
	store, _ := NewPgStoreFromConnString("postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	err := store.PutRole(shieldwall.Role{Id: "test_role", Namespaces: []string{"shieldwall.io/users", "shieldwall.io/admins"}})
	if err != nil {
		t.Error(err)
	}

	err = store.PutRole(shieldwall.Role{Id: "test_role", Namespaces: []string{"shieldwall.io/users", "shieldwall.io/admins", "shieldwall.io/companies/*/users"}})
	if err != nil {
		t.Error(err)
	}

	rows, err := store.db.Query("select * from role where id = $1", "test_role")
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
	assert.Equal(t, namespaces, []string{"shieldwall.io/users", "shieldwall.io/admins", "shieldwall.io/companies/*/users"}, "it should return namespaces as 'shieldwall.io/users', 'shieldwall.io/admins' and 'shieldwall.io/companies/*/users'")
}
