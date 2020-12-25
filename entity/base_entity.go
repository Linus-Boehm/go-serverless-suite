package entity

import (
	"fmt"

	"github.com/Linus-Boehm/go-serverless-suite/common"
)

const (
	UserEntityKey   EntityKey = "USER"
	TenantEntityKey EntityKey = "TENANT"
)

type EntityKey string

func (e EntityKey) String() string {
	return string(e)
}

type TableIndex struct {
	Name string
	PK   string
	SK   string
}

type DBKeyer interface {
	GetPK() fmt.Stringer
	GetSK() fmt.Stringer
	GetEntity() fmt.Stringer
}

type DBProvider interface {
	DBKeyer
	GetPayload() interface{}
}

// DBEntityGetter is used for tests to test all entities in one table test
type DBEntityGetter interface {
	GetDBEntity() DBProvider
}

// every GetEntity struct that should get persisted should include BaseEntity
type BaseEntity struct {
	PK     string `dynamo:"pk,hash" index:"gsi-1-reverse,range"`
	SK     string `dynamo:"sk,range" index:"gsi-1-reverse,hash"`
	Entity string `dynamo:"entity,omitempty" index:"gsi-2-entity,hash"`
	Slug   string `dynamo:"slug,omitempty" index:"gsi-2-entity,range"`
}

// BaseEntity fulfills itf.TableKey
func NewBaseEntity(pk, sk, slug, entity string) *BaseEntity {
	return &BaseEntity{
		PK:     pk,
		SK:     sk,
		Entity: entity,
		Slug:   slug,
	}
}

func (e *BaseEntity) GetPK() fmt.Stringer {
	return common.NewString(e.PK)
}

func (e *BaseEntity) GetSK() fmt.Stringer {
	return common.NewString(e.SK)
}

func (e *BaseEntity) GetEntity() fmt.Stringer {
	return UserEntityKey
}

func (e *BaseEntity) GetPayload() interface{} {
	panic("this method should be implemented in the GetEntity")
}
