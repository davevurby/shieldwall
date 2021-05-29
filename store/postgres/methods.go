package postgres

import (
	"log"

	"github.com/davevurby/shieldwall"
	"github.com/lib/pq"
)

func (s *PgStore) GetRoles() ([]shieldwall.Role, error) {
	return nil, nil
}

func (s *PgStore) PutRole(role shieldwall.Role) error {
	log.Printf("Upserting role %s...\n", role.Id)

	_, err := s.db.Exec("insert into role (id, namespaces) values ($1, $2) on conflict (id) do update set namespaces = $2", role.Id, pq.Array(role.Namespaces))
	return err
}

func (s *PgStore) CreateIdentity(identity shieldwall.Identity) error {
	log.Printf("Upserting identity %s...\n", identity.Id)

	_, err := s.db.Exec("insert into identity (id, namespace, roles) values ($1, $2, $3) on conflict (id, namespace) do update set roles = $3", identity.Id, identity.Namespace, pq.Array(identity.Roles))
	return err
}

func (s *PgStore) CreatePolicy(policy shieldwall.Policy) error {
	log.Printf("Upserting policy...\n")

	_, err := s.db.Exec("insert into policy (subject, namespace, policy, effect) values ($1, $2, $3, $4) on conflict (subject, namespace, policy, effect) do nothing", policy.Subject, policy.Namespace, policy.Object, policy.Effect)
	return err
}

func (s *PgStore) IsPermitted(subject string, namespace string, object string, effect string) (bool, error) {
	query := `
		select count(p.*) from policy p
		left join role r on r.id = p.subject
		left join identity i on r.id = any(i.roles)
		where
		 $1 ~ p.namespace
		   and p.policy = $2
		   and p.effect = $3
		   and (i.id = $4 or p.subject = $4)
	`

	rows, err := s.db.Query(query, namespace, object, effect, subject)
	if err != nil {
		return false, err
	}

	rows.Next()
	var count int
	if err = rows.Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}
