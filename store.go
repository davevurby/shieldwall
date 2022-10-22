package shieldwall

type Store interface {
	// Close closes a connection to a store
	Close() error

	GetRole(id string) (*Role, error)
	GetRoles() ([]Role, error)
	PutRole(role Role) error
	DeleteRole(id string) error

	GetIdentity(id string) (*Identity, error)
	GetIdentities() ([]Identity, error)
	PutIdentity(identity Identity) error
	DeleteIdentity(id string) error

	PutPolicy(policy Policy) error
	IsPermitted(subject string, namespace string, object string, effect string) (bool, error)
}
