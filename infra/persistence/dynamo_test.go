package persistence

import (
	"testing"

	"github.com/Linus-Boehm/go-serverless-suite/repositories"

	"github.com/Linus-Boehm/go-serverless-suite/entity"
	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	b, err := NewTestProvider(repositories.UserEntity{})
	assert.NoError(t, err)
	assert.NoError(t, b.DeleteTable())
}

func TestDynamoBaseTable_PutReadItem(t *testing.T) {
	b, err := NewTestProvider(repositories.UserEntity{})
	assert.NoError(t, err)
	defer b.DeleteTable()
	ts := entity.Timestamps{}
	ts.CreatedNow()
	user := entity.User{
		ID:         entity.NewEntityIDV4(),
		Email:      "example.org",
		Firstname:  "Linus",
		Lastname:   "Boehm",
		Attributes: map[string]string{},
		Timestamps: ts,
	}
	e := repositories.NewUserEntity(user)
	err = b.PutItem(e)
	assert.NoError(t, err)
	oldU := e.(*repositories.UserEntity)
	newU := repositories.UserEntity{}
	err = b.ReadItem(e, &newU)
	assert.NoError(t, err)

	assert.Equal(t, oldU.PK, newU.PK)
	assert.Equal(t, oldU.SK, newU.SK)
	assert.Equal(t, oldU.Firstname, newU.Firstname)
	assert.Equal(t, oldU.Lastname, newU.Lastname)
}

func TestDynamoBaseTable_ReadAllWithPK(t *testing.T) {
	b, err := NewTestProvider(repositories.UserEntity{})
	assert.NoError(t, err)
	defer b.DeleteTable()
	ts := entity.Timestamps{}
	ts.CreatedNow()
	user := entity.User{
		ID:         entity.NewEntityIDV4(),
		Email:      "example.org",
		Firstname:  "Linus",
		Lastname:   "Boehm",
		Attributes: map[string]string{},
		Timestamps: ts,
	}
	e := repositories.NewUserEntity(user)
	err = b.PutItem(e)
	assert.NoError(t, err)
	result := []repositories.UserEntity{}
	err = b.ReadAllWithPK(e.GetPK(), nil, entity.UserEntityName, &result)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, result[0].PK, e.GetPK().String())
}

func TestDynamoBaseTable_BatchWriteItems(t *testing.T) {
	b, err := NewTestProvider(repositories.UserEntity{})
	assert.NoError(t, err)
	defer b.DeleteTable()
	newU := func() entity.User {
		return entity.User{
			ID:    entity.NewEntityIDV4(),
			Email: entity.NewEntityIDV4().String(),
		}
	}
	var payload []interface{}

	for i := 0; i < 50; i++ {
		u := repositories.NewUserEntity(newU())
		payload = append(payload, &u)
	}

	err = b.BatchWriteItems(payload...)
	assert.NoError(t, err)
	var counter []repositories.UserEntity
	err = b.GetEntity(DynamoEntityIndex, entity.UserEntityName, &counter, false)
	assert.NoError(t, err)
	assert.Len(t, counter, 50)
}
