package entity

import (
	"fmt"

	"github.com/Linus-Boehm/go-serverless-suite/common"
)

const (
	TenantEntityName string = "TENANT"
)

type Tenant struct {
	ID       ID
	ParentID ID
	Name     string
	Slug     string
	Timestamps
}

func NewTenant(name, slug string, parent ID) Tenant {
	id := NewEntityIDV4()
	if parent.IsNil() {
		parent = id
	}
	t := Tenant{
		ID:       id,
		ParentID: parent,
		Name:     name,
		Slug:     slug,
	}

	t.CreatedNow()
	return t
}

type TenantEntity struct {
	BaseEntity
	Timestamps Timestamps `dynamo:"timestamps,omitempty"`
	Name       string     `dynamo:"name,omitempty"`
}

func (t Tenant) GetDBEntity() DBProvider {
	return &TenantEntity{
		BaseEntity: BaseEntity{
			PK:     common.JoinDBKey(TenantEntityName, t.ParentID.String()),
			SK:     common.JoinDBKey(TenantEntityName, t.ID.String()),
			Entity: TenantEntityName,
			Slug:   t.Slug,
		},
		Timestamps: t.Timestamps,
		Name:       t.Name,
	}
}

func (e *TenantEntity) GetTenant() (Tenant, error) {
	id, err := IDFromDBString(e.SK)
	if err != nil {
		return Tenant{}, err
	}
	parent, err := IDFromDBString(e.PK)
	if err != nil {
		return Tenant{}, err
	}
	return Tenant{
		ID:         id,
		ParentID:   parent,
		Name:       e.Name,
		Slug:       e.Slug,
		Timestamps: e.Timestamps,
	}, nil
}

func (e *TenantEntity) GetTimestamps() Timestamps {
	return e.Timestamps
}

func (e *TenantEntity) GetPK() fmt.Stringer {
	return common.NewString(e.PK)
}

func (e *TenantEntity) GetSK() fmt.Stringer {
	return common.NewString(e.SK)
}

func (e *TenantEntity) GetEntity() fmt.Stringer {
	return common.NewString(TenantEntityName)
}

func (e *TenantEntity) GetPayload() interface{} {
	return e
}
