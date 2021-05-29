package postgres

import (
	"log"

	"github.com/davevurby/shieldwall"
	"github.com/lib/pq"
)

func (s *PgStore) GetRole(id string) (*shieldwall.Role, error) {
	query := `
		select r.id, r.namespaces from role r where r.id = $1
	`

	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	rows.Next()
	var role shieldwall.Role
	if err = rows.Scan(&role.Id, pq.Array(&role.Namespaces)); err != nil {
		return nil, err
	}

	return &role, nil
}

func (s *PgStore) GetRoles() ([]shieldwall.Role, error) {
	query := `
		select r.id, r.namespaces from role r
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	var roles []shieldwall.Role

	for rows.Next() {
		var role shieldwall.Role

		if err := rows.Scan(&role.Id, pq.Array(&role.Namespaces)); err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	return roles, nil
}

func (s *PgStore) PutRole(role shieldwall.Role) error {
	log.Printf("Upserting role %s...\n", role.Id)

	_, err := s.db.Exec("insert into role (id, namespaces) values ($1, $2) on conflict (id) do update set namespaces = $2", role.Id, pq.Array(role.Namespaces))
	return err
}
