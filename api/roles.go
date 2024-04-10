package api

var (
	AdminRole = &accessRole{name: "admin"}
	UserRole  = &accessRole{name: "user"}
)

type AccessRole interface {
	GetName() string
}

type accessRole struct {
	name string
}

func (ar *accessRole) GetName() string {
	return ar.name
}
