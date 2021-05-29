package shieldwall

type Identity struct {
	Id        string   `json:"id"`
	Namespace string   `json:"namespaces"`
	Roles     []string `json:"roles"`
}

type Role struct {
	Id         string   `json:"id"`
	Namespaces []string `json:"namespaces"`
}

type Policy struct {
	Subject   string `json:"id"`
	Namespace string `json:"namespaces"`
	Object    string `json:"object"`
	Effect    string `json:"effect"`
}
