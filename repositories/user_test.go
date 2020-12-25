package repositories

import (
	"testing"

	"github.com/Linus-Boehm/go-serverless-suite/entity"
	"github.com/Linus-Boehm/go-serverless-suite/infra/persistence"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_ReadUser(t *testing.T) {
	b, err := persistence.NewTestProvider(entity.UserEntity{})
	assert.NoError(t, err)
	defer b.DeleteTable()

	userP := NewUserRepository(b)

	_, err = userP.ReadUser(entity.NewEntityIDV4())
	assert.Error(t, err)

	tenant := entity.Tenant{
		ID:       entity.NewEntityIDV4(),
		ParentID: entity.NewEntityIDV4(),
		Name:     "tenant",
		Slug:     "tenant",
	}

	err = b.PutItem(tenant.GetDBEntity())
	assert.NoError(t, err)
	_, err = userP.ReadUser(tenant.ID)
	assert.Error(t, err)

	user := entity.User{
		ID:    entity.NewEntityIDV4(),
		Email: "email@example.org",
	}
	err = b.PutItem(user.GetDBEntity())
	assert.NoError(t, err)

	newUser, err := userP.ReadUser(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, newUser.Email)
}
