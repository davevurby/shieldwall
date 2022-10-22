package postgres

import (
	"log"

	"github.com/davevurby/shieldwall"
	"github.com/lib/pq"
)

func (s *PgStore) GetIdentity(id string) (*shieldwall.Identity, error) {
	query := `
		select i.id, i.namespace, i.roles from identity i where i.id = $1
	`

	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, nil
	}

	var identity shieldwall.Identity
	if err = rows.Scan(&identity.Id, &identity.Namespace, pq.Array(&identity.Roles)); err != nil {
		return nil, err
	}

	return &identity, nil
}

func (s *PgStore) GetIdentities() ([]shieldwall.Identity, error) {
	query := `
		select i.id, i.namespace, i.roles from identity i
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	var identities []shieldwall.Identity

	for rows.Next() {
		var identity shieldwall.Identity

		if err := rows.Scan(&identity.Id, &identity.Namespace, pq.Array(&identity.Roles)); err != nil {
			return nil, err
		}

		identities = append(identities, identity)
	}

	return identities, nil
}

func (s *PgStore) PutIdentity(identity shieldwall.Identity) error {
	log.Printf("Upserting identity %s...\n", identity.Id)

	_, err := s.db.Exec("insert into identity (id, namespace, roles) values ($1, $2, $3) on conflict (id, namespace) do update set roles = $3", identity.Id, identity.Namespace, pq.Array(identity.Roles))
	return err
}

func (s *PgStore) DeleteIdentity(id string) error {
	_, err := s.db.Exec("delete from identity where id = $1", id)
	return err
}
