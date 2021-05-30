package postgres

import (
	"log"

	"github.com/davevurby/shieldwall"
)

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
