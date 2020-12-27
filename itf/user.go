package itf

import "github.com/Linus-Boehm/go-serverless-suite/entity"

type TenantProvider interface {
}

type TenantServicer interface {
}

type UserServicer interface {
}

type UserProvider interface {
	ReadUser(id entity.ID) (entity.User, error)
	ReadUserByEmail(email entity.ID) (entity.User, error)
	PutUser(user entity.User) error
	DeleteUser(id entity.ID, email string, soft bool) (entity.User, error)
	ListUsers() ([]entity.User, error)
}
