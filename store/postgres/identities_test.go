package postgres

import (
	"testing"

	"github.com/davevurby/shieldwall"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestPgStore_PutIdentity(t *testing.T) {
	store, _ := NewPgStoreFromConnString("postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	err := store.PutIdentity(shieldwall.Identity{Id: "johndoe", Namespace: "shieldwall.io/users", Roles: []string{"shieldwall.io/roles/foo", "shieldwall.io/roles/bar"}})
	if err != nil {
		t.Error(err)
	}

	rows, err := store.db.Query("select id, namespace, roles from identity where id = $1", "johndoe")
	if err != nil {
		t.Error(err)
	}

	result := rows.Next()
	if result == false {
		t.Error("there is no identity")
	}

	var id string
	var namespace string
	var roles []string
	if err = rows.Scan(&id, &namespace, pq.Array(&roles)); err != nil {
		t.Error(err)
	}

	result = rows.Next()
	if result == true {
		t.Error("there should be no more identity")
	}

	assert.Equal(t, id, "johndoe", "it should return id as 'id'")
	assert.Equal(t, namespace, "shieldwall.io/users", "it should return correct namespace")
	assert.Equal(t, roles, []string{"shieldwall.io/roles/foo", "shieldwall.io/roles/bar"}, "it should return roles as 'shieldwall.io/roles/foo' and 'shieldwall.io/roles/bar'")
}

func TestPgStore_GetIdentity(t *testing.T) {
	store, _ := NewPgStoreFromConnString("postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	err := store.PutIdentity(shieldwall.Identity{Id: "johndoe", Namespace: "shieldwall.io/users", Roles: []string{"shieldwall.io/roles/foo", "shieldwall.io/roles/bar"}})
	if err != nil {
		t.Error(err)
	}

	identity, err := store.GetIdentity("johndoe")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, identity.Id, "johndoe", "it should return id as 'johndoe'")
	assert.Equal(t, identity.Namespace, "shieldwall.io/users", "it should return correct namespace")
	assert.Equal(t, identity.Roles, []string{"shieldwall.io/roles/foo", "shieldwall.io/roles/bar"}, "it should return roles as 'shieldwall.io/roles/foo' and 'shieldwall.io/roles/bar'")
}

func TestPgStore_GetIdentities(t *testing.T) {
	store, _ := NewPgStoreFromConnString("postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	err := store.PutIdentity(shieldwall.Identity{Id: "johndoe", Namespace: "shieldwall.io/users", Roles: []string{"shieldwall.io/roles/foo", "shieldwall.io/roles/bar"}})
	if err != nil {
		t.Error(err)
	}

	identities, err := store.GetIdentities()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, identities[0].Id, "johndoe", "it should return id as 'johndoe'")
	assert.Equal(t, identities[0].Namespace, "shieldwall.io/users", "it should return correct namespace")
	assert.Equal(t, identities[0].Roles, []string{"shieldwall.io/roles/foo", "shieldwall.io/roles/bar"}, "it should return roles as 'shieldwall.io/roles/foo' and 'shieldwall.io/roles/bar'")
}

func TestPgStore_DeleteIdentity(t *testing.T) {
	store, _ := NewPgStoreFromConnString("postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	err := store.PutIdentity(shieldwall.Identity{Id: "johndoe", Namespace: "shieldwall.io/users", Roles: []string{"shieldwall.io/roles/foo", "shieldwall.io/roles/bar"}})
	if err != nil {
		t.Error(err)
	}

	err = store.DeleteIdentity("johndoe")
	if err != nil {
		t.Error(err)
	}

	identity, err := store.GetIdentity("johndoe")
	if err != nil {
		t.Error(err)
	}

	if identity != nil {
		t.Error("identity should be nil")
	}
}
