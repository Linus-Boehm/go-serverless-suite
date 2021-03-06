package itf

import "github.com/Linus-Boehm/go-serverless-suite/entity"

//go:generate mockgen -destination=user_mocks.go -package=itf -source=user.go

type TenantProvider interface {
}

type TenantServicer interface {
}

type UserServicer interface {
	CreateNewUser(user entity.User) (entity.User, error)
}

type UserProvider interface {
	ReadUser(id entity.ID) (entity.User, error)
	ReadUserByEmail(email entity.ID) (entity.User, error)
	PutUser(user entity.User) error
	DeleteUser(id entity.ID, email string, soft bool) (entity.User, error)
	ListUsers() ([]entity.User, error)
}
