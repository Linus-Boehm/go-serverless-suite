package repositories

import (
	"fmt"

	"github.com/Linus-Boehm/go-serverless-suite/entity"

	"github.com/Linus-Boehm/go-serverless-suite/common"
)

type TableIndex struct {
	Name string
	PK   string
	SK   string
}

// every GetEntity struct that should get persisted should include BaseEntity
type BaseEntity struct {
	PK         string            `dynamo:"pk,hash" index:"gsi-1-reverse,range"`
	SK         string            `dynamo:"sk,range" index:"gsi-1-reverse,hash"`
	Entity     string            `dynamo:"entity,omitempty" index:"gsi-2-entity,hash"`
	Slug       string            `dynamo:"slug,omitempty" index:"gsi-2-entity,range"`
	Timestamps entity.Timestamps `dynamo:"timestamps,omitempty"`
}

// BaseEntity fulfills itf.TableKey
func NewBaseEntity(pkID entity.ID, skID entity.ID, slug string, entity entity.Name) BaseEntity {
	return BaseEntity{
		PK:     common.JoinStringerDBKey(entity, pkID),
		SK:     common.JoinStringerDBKey(entity, skID),
		Entity: entity.String(),
		Slug:   slug,
	}
}

func (e BaseEntity) WithTimestamps(t entity.Timestamps) BaseEntity {
	e.Timestamps = t
	return e
}

func (e *BaseEntity) GetPK() fmt.Stringer {
	return common.NewString(e.PK)
}

func (e *BaseEntity) GetSK() fmt.Stringer {
	return common.NewString(e.SK)
}

func (e *BaseEntity) GetEntity() fmt.Stringer {
	return common.NewString(e.Entity)
}
