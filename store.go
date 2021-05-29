package shieldwall

type Store interface {
	// Close closes a connection to a store
	Close() error

	GetRoles() ([]Role, error)
	PutRole(role Role) error

	CreateIdentity(identity Identity) error
	CreatePolicy(policy Policy) error
	IsPermitted(subject string, namespace string, object string, effect string) (bool, error)
}
