package entity

import (
	"fmt"

	"github.com/Linus-Boehm/go-serverless-suite/common"
)

type User struct {
	ID         ID
	Email      string
	Firstname  string
	Lastname   string
	Attributes map[string]string
	Timestamps
}

// UserEntity is implementing itf.TableEntity
type UserEntity struct {
	BaseEntity
	Timestamps     Timestamps        `dynamo:"timestamps,omitempty"`
	Firstname      string            `dynamo:"firstname,omitempty"`
	Lastname       string            `dynamo:"lastname,omitempty"`
	UserAttributes map[string]string `dynamo:"user_attributes,omitempty"`
}

func (u User) GetDBEntity() DBProvider {
	return &UserEntity{
		BaseEntity: BaseEntity{
			PK:     common.JoinStringerDBKey(UserEntityKey, u.ID),
			SK:     common.JoinDBKey(UserEntityKey.String(), u.Email),
			Entity: UserEntityKey.String(),
			Slug:   fmt.Sprintf("user-%s", u.ID.String()),
		},
		Timestamps:     u.Timestamps,
		Firstname:      u.Lastname,
		Lastname:       u.Lastname,
		UserAttributes: u.Attributes,
	}
}

func (e *UserEntity) GetUser() (User, error) {
	id, err := IDFromDBString(e.PK)
	if err != nil {
		return User{}, err
	}
	email, err := IDFromDBString(e.SK)
	if err != nil {
		return User{}, err
	}
	u := User{
		ID:         id,
		Email:      email.String(),
		Firstname:  e.Firstname,
		Lastname:   e.Lastname,
		Attributes: e.UserAttributes,
		Timestamps: e.Timestamps,
	}

	return u, nil
}

func (e *UserEntity) GetTimestamps() Timestamps {
	return e.Timestamps
}

func (e *UserEntity) GetPK() fmt.Stringer {
	return common.NewString(e.PK)
}

func (e *UserEntity) GetSK() fmt.Stringer {
	return common.NewString(e.SK)
}

func (e *UserEntity) GetEntity() fmt.Stringer {
	return UserEntityKey
}

func (e *UserEntity) GetPayload() interface{} {
	return e
}
