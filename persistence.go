package mama_keeper

type MamaKeeperPersistence interface {
	CreateIdentity(identity Identity) error
	CreateRole(role Role) error
	CreatePolicy(policy Policy) error
	IsPermitted(subject string, namespace string, object string, effect string) (bool, error)
	Close() error
}
