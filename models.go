package mamakeeper

type Identity struct {
	Id        string
	Namespace string
	Roles     []string
}

type Role struct {
	Id         string
	Namespaces []string
}

type Policy struct {
	Subject   string
	Namespace string
	Object    string
	Effect    string
}
