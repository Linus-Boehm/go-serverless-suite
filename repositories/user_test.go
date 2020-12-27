package repositories

import (
	"testing"
	"time"

	"github.com/Linus-Boehm/go-serverless-suite/itf"

	"github.com/Linus-Boehm/go-serverless-suite/entity"
	"github.com/Linus-Boehm/go-serverless-suite/infra/persistence"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_ReadUser(t *testing.T) {
	now := time.Now().Unix()

	tests := []struct {
		name       string
		inputID    entity.ID
		beforeTest func(t *testing.T, b itf.BaseTableProvider, id entity.ID)
		expectUser func(id entity.ID) entity.User
		expectErr  bool
	}{
		{
			name:    "happy - find user",
			inputID: entity.NewEntityIDV4(),
			beforeTest: func(t *testing.T, b itf.BaseTableProvider, id entity.ID) {
				user := entity.User{
					ID:    id,
					Email: "email@example.org",
				}
				err := b.PutItem(NewUserEntity(user))
				assert.NoError(t, err)
			},
			expectUser: func(id entity.ID) entity.User {
				return entity.User{
					ID:    id,
					Email: "email@example.org",
				}
			},
			expectErr: false,
		},
		{
			name:    "happy - find user wit all attr",
			inputID: entity.NewEntityIDV4(),
			beforeTest: func(t *testing.T, b itf.BaseTableProvider, id entity.ID) {
				user := entity.User{
					ID:         id,
					Email:      "email@example.org",
					Firstname:  "Max",
					Lastname:   "Muster",
					Attributes: map[string]string{"Foo": "Bar"},
					Timestamps: entity.Timestamps{
						CreatedAt:   now,
						UpdatedAt:   now,
						PublishedAt: &now,
						DeletedAt:   nil,
					},
				}
				err := b.PutItem(NewUserEntity(user))
				assert.NoError(t, err)
			},
			expectUser: func(id entity.ID) entity.User {
				return entity.User{
					ID:         id,
					Email:      "email@example.org",
					Firstname:  "Max",
					Lastname:   "Muster",
					Attributes: map[string]string{"Foo": "Bar"},
					Timestamps: entity.Timestamps{
						CreatedAt:   now,
						UpdatedAt:   now,
						PublishedAt: &now,
						DeletedAt:   nil,
					},
				}
			},
			expectErr: false,
		},
		{
			name:    "happy - don't read other entities",
			inputID: entity.NewEntityIDV4(),
			beforeTest: func(t *testing.T, b itf.BaseTableProvider, id entity.ID) {
				tenant := entity.Tenant{
					ID:       id,
					ParentID: id,
					Name:     "tenant",
					Slug:     "tenant",
				}
				err := b.PutItem(NewTenantEntity(tenant))
				assert.NoError(t, err)
			},
			expectUser: func(id entity.ID) entity.User {
				return entity.User{}
			},
			expectErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b, err := persistence.NewTestProvider(UserEntity{})
			assert.NoError(t, err)
			defer b.DeleteTable()

			repo := NewUserRepository(b)
			if test.beforeTest != nil {
				test.beforeTest(t, b, test.inputID)
			}
			u, err := repo.ReadUser(test.inputID)
			assert.Equal(t, test.expectErr, err != nil)
			assert.EqualValues(t, test.expectUser(test.inputID), u)
		})
	}
}

func TestUserRepository_DeleteUser(t *testing.T) {
	now := time.Now().Unix()
	tests := []struct {
		name       string
		inputID    entity.ID
		inputEmail string
		beforeTest func(t *testing.T, b itf.BaseTableProvider, id entity.ID, inputEmail string)
		soft       bool
		expectUser func(id entity.ID, inputEmail string) entity.User
		expectErr  bool
	}{
		{
			name:       "happy - delete user",
			inputID:    entity.NewEntityIDV4(),
			inputEmail: "test@example.org",
			beforeTest: func(t *testing.T, b itf.BaseTableProvider, id entity.ID, inputEmail string) {
				user := entity.User{
					ID:        id,
					Email:     inputEmail,
					Firstname: "Max",
				}
				err := b.PutItem(NewUserEntity(user))
				assert.NoError(t, err)
			},
			expectUser: func(id entity.ID, inputEmail string) entity.User {
				return entity.User{
					ID:        id,
					Email:     inputEmail,
					Firstname: "Max",
				}
			},
			expectErr: false,
		},
		{
			name:       "happy - delete user soft",
			inputID:    entity.NewEntityIDV4(),
			inputEmail: "test@example.org",
			beforeTest: func(t *testing.T, b itf.BaseTableProvider, id entity.ID, inputEmail string) {
				d := int64(0)
				user := entity.User{
					ID:        id,
					Email:     inputEmail,
					Firstname: "Max",
					Timestamps: entity.Timestamps{
						CreatedAt:   now,
						UpdatedAt:   now,
						PublishedAt: nil,
						DeletedAt:   &d,
					},
				}
				err := b.PutItem(NewUserEntity(user))
				assert.NoError(t, err)
			},
			soft: true,
			expectUser: func(id entity.ID, inputEmail string) entity.User {
				return entity.User{
					ID:        id,
					Email:     inputEmail,
					Firstname: "Max",
				}
			},
			expectErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b, err := persistence.NewTestProvider(UserEntity{})
			assert.NoError(t, err)
			defer b.DeleteTable()

			repo := NewUserRepository(b)
			if test.beforeTest != nil {
				test.beforeTest(t, b, test.inputID, test.inputEmail)
			}
			u, err := repo.DeleteUser(test.inputID, test.inputEmail, test.soft)
			assert.Equal(t, test.expectErr, err != nil)
			if test.soft {
				assert.NotNil(t, u)
				assert.NotNil(t, u.Timestamps.DeletedAt)
				assert.Greater(t, *u.Timestamps.DeletedAt, int64(0))
				u.Timestamps = entity.Timestamps{}
			}
			assert.EqualValues(t, test.expectUser(test.inputID, test.inputEmail), u)
		})
	}
}

func TestUserRepository_PutUser(t *testing.T) {
	now := time.Now().Unix()

	tests := []struct {
		name       string
		inputID    entity.ID
		beforeTest func(t *testing.T, b itf.BaseTableProvider, id entity.ID)
		putUser    func(id entity.ID) entity.User
		expectUser func(id entity.ID) entity.User
		expectErr  bool
	}{
		{
			name:    "happy - create user",
			inputID: entity.NewEntityIDV4(),
			putUser: func(id entity.ID) entity.User {
				return entity.User{
					ID:        id,
					Email:     "email@example.org",
					Firstname: "Max",
					Lastname:  "Muster",
					Timestamps: entity.Timestamps{
						CreatedAt:   now,
						UpdatedAt:   now,
						PublishedAt: nil,
						DeletedAt:   nil,
					},
				}
			},
			expectUser: func(id entity.ID) entity.User {
				return entity.User{
					ID:        id,
					Email:     "email@example.org",
					Firstname: "Max",
					Lastname:  "Muster",
					Timestamps: entity.Timestamps{
						CreatedAt:   now,
						UpdatedAt:   now,
						PublishedAt: nil,
						DeletedAt:   nil,
					},
				}
			},
			expectErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b, err := persistence.NewTestProvider(UserEntity{})
			assert.NoError(t, err)
			defer b.DeleteTable()

			repo := NewUserRepository(b)
			if test.beforeTest != nil {
				test.beforeTest(t, b, test.inputID)
			}
			err = repo.PutUser(test.putUser(test.inputID))
			assert.Equal(t, test.expectErr, err != nil)
			if !test.expectErr {
				user, err := repo.ReadUser(test.inputID)
				assert.NoError(t, err)
				assert.EqualValues(t, test.expectUser(test.inputID), user)
			}

		})
	}
}
