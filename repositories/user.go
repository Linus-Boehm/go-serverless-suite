package repositories

import (
	"strings"

	"github.com/Linus-Boehm/go-serverless-suite/entity"
	"github.com/Linus-Boehm/go-serverless-suite/itf"
	"github.com/pkg/errors"
)

type userRepository struct {
	table itf.BaseTableProvider
}

func NewUserRepository(basetable itf.BaseTableProvider) itf.UserProvider {
	return &userRepository{table: basetable}
}

func (u userRepository) ReadUser(id entity.ID) (user entity.User, err error) {
	rows := []entity.UserEntity{}
	err = u.table.ReadAllWithPK(id.WithEntity(entity.UserEntityKey), nil, entity.UserEntityKey, &rows)
	if err != nil {
		return entity.User{}, err
	}
	for _, row := range rows {
		if strings.HasPrefix(row.SK, entity.UserEntityKey.String()) {
			u, err := row.GetUser()
			if err != nil {
				return user, err
			}
			return u, nil
		}
	}
	return entity.User{}, errors.New("could not find a convertible user entity")
}

func (u userRepository) ReadUserByEmail(email entity.ID) (entity.User, error) {
	panic("implement me")
}

func (u userRepository) PutUser(user entity.User) error {
	panic("implement me")
}

func (u userRepository) DeleteUser(id entity.ID, email string, soft bool) error {
	panic("implement me")
}

func (u userRepository) ListUsers() ([]entity.User, error) {
	panic("implement me")
}
