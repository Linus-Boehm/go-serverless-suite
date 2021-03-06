package repositories

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/Linus-Boehm/go-serverless-suite/infra/persistence"

	"github.com/Linus-Boehm/go-serverless-suite/common"

	"github.com/Linus-Boehm/go-serverless-suite/entity"
	"github.com/Linus-Boehm/go-serverless-suite/itf"
)

type userRepository struct {
	table itf.BaseTableProvider
}

func NewUserRepository(basetable itf.BaseTableProvider) itf.UserProvider {
	return &userRepository{table: basetable}
}

func (u userRepository) ReadUser(id entity.ID) (user entity.User, err error) {
	rows := []UserEntity{}
	err = u.table.ReadAllWithPK(id.WithEntity(entity.UserEntityName), nil, entity.UserEntityName, &rows)
	if err != nil {
		return entity.User{}, err
	}
	for _, row := range rows {
		if strings.HasPrefix(row.SK, entity.UserEntityName.String()) {
			u, err := row.GetUser()
			if err != nil {
				return user, err
			}
			return u, nil
		}
	}
	return entity.User{}, common.NewEntityNotFoundError(id, entity.UserEntityName)
}

func (u userRepository) ReadUserByEmail(email entity.ID) (user entity.User, err error) {
	rows := []UserEntity{}
	err = u.table.ReadAllWithPK(email.WithEntity(entity.UserEntityName), &persistence.DynamoReverseIndex, entity.UserEntityName, &rows)
	if err != nil {
		return entity.User{}, err
	}
	for _, row := range rows {
		if strings.HasPrefix(row.SK, entity.UserEntityName.String()) {
			u, err := row.GetUser()
			if err != nil {
				return user, err
			}
			return u, nil
		}
	}
	return user, common.NewEntityNotFoundError(email, entity.UserEntityName)
}

func (u userRepository) PutUser(user entity.User) error {
	return u.table.PutItem(NewUserEntity(user))
}

func (u userRepository) DeleteUser(id entity.ID, email string, soft bool) (entity.User, error) {
	user := UserEntity{}
	key := common.NewStringKey(id.String(), email, entity.UserEntityName)

	if soft {
		err := u.table.RemoveItemSoft(key, &user)
		if err != nil {
			return entity.User{}, err
		}
		return user.GetUser()
	}
	err := u.table.RemoveItem(key, &user)
	if err != nil {
		return entity.User{}, err
	}
	return user.GetUser()
}

func (u userRepository) ListUsers() (users []entity.User, err error) {
	rows := []UserEntity{}

	err = u.table.GetEntity(persistence.DynamoEntityIndex, entity.UserEntityName, &rows, true)
	if err != nil {
		return nil, err
	}
	return mapUsersToDomain(rows)
}

func mapUsersToDomain(users []UserEntity) ([]entity.User, error) {
	u := []entity.User{}
	for _, user := range users {
		e, err := user.GetUser()
		if err != nil {
			return nil, errors.Wrap(err, "unable to transfer user row to entity")
		}
		u = append(u, e)
	}
	return u, nil
}

// UserEntity is implementing itf.TableEntity
type UserEntity struct {
	BaseEntity
	Firstname      string            `dynamo:"firstname,omitempty"`
	Lastname       string            `dynamo:"lastname,omitempty"`
	EmailVerified  bool              `dynamo:"emailVerified,omitempty"`
	UserAttributes map[string]string `dynamo:"user_attributes,omitempty"`
}

func NewUserEntity(u entity.User) *UserEntity {
	return &UserEntity{
		BaseEntity: BaseEntity{
			PK:         common.JoinStringerDBKey(entity.UserEntityName, u.ID),
			SK:         common.JoinDBKey(entity.UserEntityName.String(), u.Email),
			Entity:     entity.UserEntityName.String(),
			Slug:       fmt.Sprintf("user-%s", u.ID.String()),
			Timestamps: u.Timestamps,
		},
		EmailVerified:  u.EmailVerified,
		Firstname:      u.Firstname,
		Lastname:       u.Lastname,
		UserAttributes: u.Attributes,
	}
}

func (e *UserEntity) GetUser() (entity.User, error) {
	id, err := entity.IDFromDBString(e.PK)
	if err != nil {
		return entity.User{}, err
	}
	email, err := entity.IDFromDBString(e.SK)
	if err != nil {
		return entity.User{}, err
	}
	u := entity.User{
		ID:            id,
		Email:         email.String(),
		Firstname:     e.Firstname,
		Lastname:      e.Lastname,
		Attributes:    e.UserAttributes,
		Timestamps:    e.Timestamps,
		EmailVerified: e.EmailVerified,
	}

	return u, nil
}

func (e *UserEntity) GetTimestamps() entity.Timestamps {
	return e.Timestamps
}

func (e *UserEntity) GetPK() fmt.Stringer {
	return common.NewString(e.PK)
}

func (e *UserEntity) GetSK() fmt.Stringer {
	return common.NewString(e.SK)
}

func (e *UserEntity) GetEntity() fmt.Stringer {
	return entity.UserEntityName
}

func (e *UserEntity) IsDeleted() bool {
	return e.Timestamps.IsDeleted()
}

func (e *UserEntity) SoftDeleteNow() {
	e.Timestamps.SoftDeleteNow()
}
