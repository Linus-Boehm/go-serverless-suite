package repositories

import (
	"fmt"

	"github.com/Linus-Boehm/go-serverless-suite/common"
	"github.com/Linus-Boehm/go-serverless-suite/entity"
	"github.com/Linus-Boehm/go-serverless-suite/itf"
)

type TenantEntity struct {
	BaseEntity
	Name string `dynamo:"name,omitempty"`
}

func NewTenantEntity(t entity.Tenant) itf.DBKeyer {
	return &TenantEntity{
		BaseEntity: NewBaseEntity(t.ParentID, t.ID, t.Slug, entity.TenantEntityName).WithTimestamps(t.Timestamps),
		Name:       t.Name,
	}
}

func (e *TenantEntity) GetTenant() (entity.Tenant, error) {
	id, err := entity.IDFromDBString(e.SK)
	if err != nil {
		return entity.Tenant{}, err
	}
	parent, err := entity.IDFromDBString(e.PK)
	if err != nil {
		return entity.Tenant{}, err
	}
	return entity.Tenant{
		ID:         id,
		ParentID:   parent,
		Name:       e.Name,
		Slug:       e.Slug,
		Timestamps: e.Timestamps,
	}, nil
}

func (e *TenantEntity) GetTimestamps() entity.Timestamps {
	return e.Timestamps
}

func (e *TenantEntity) GetPK() fmt.Stringer {
	return common.NewString(e.PK)
}

func (e *TenantEntity) GetSK() fmt.Stringer {
	return common.NewString(e.SK)
}

func (e *TenantEntity) GetEntity() fmt.Stringer {
	return entity.TenantEntityName
}
