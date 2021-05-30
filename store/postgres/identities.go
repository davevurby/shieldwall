package postgres

import (
	"log"

	"github.com/davevurby/shieldwall"
	"github.com/lib/pq"
)

func (s *PgStore) PutIdentity(identity shieldwall.Identity) error {
	log.Printf("Upserting identity %s...\n", identity.Id)

	_, err := s.db.Exec("insert into identity (id, namespace, roles) values ($1, $2, $3) on conflict (id, namespace) do update set roles = $3", identity.Id, identity.Namespace, pq.Array(identity.Roles))
	return err
}
