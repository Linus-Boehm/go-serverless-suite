package repositories

import (
	"fmt"
	"testing"

	"github.com/Linus-Boehm/go-serverless-suite/entity"
	"github.com/Linus-Boehm/go-serverless-suite/itf"
	"github.com/stretchr/testify/assert"
)

func TestEntity(t *testing.T) {
	uid := entity.NewEntityIDV4()

	mainTenant := entity.NewTenant("Foo", "foo", entity.ID{})
	subTenant := entity.NewTenant("Bar", "bar", mainTenant.ID)

	tests := []struct {
		name         string
		entity       func() itf.DBKeyer
		expectEntity string
		expectPK     string
		expectSK     string
	}{
		{
			name: "Happy - User with uuid",
			entity: func() itf.DBKeyer {
				u := entity.User{
					ID:         uid,
					Email:      "example@org.com",
					Firstname:  "Linus",
					Lastname:   "Boehm",
					Attributes: nil,
				}
				return NewUserEntity(u)

			},
			expectPK:     fmt.Sprintf("USER#%s", uid.String()),
			expectSK:     "USER#example@org.com",
			expectEntity: "USER",
		},
		{
			name: "Happy - User with string id",
			entity: func() itf.DBKeyer {
				u := entity.User{
					ID:         entity.IDFromStringOrNil("username"),
					Email:      "example@org.com",
					Firstname:  "Linus",
					Lastname:   "Boehm",
					Attributes: nil,
				}
				return NewUserEntity(u)

			},
			expectPK:     "USER#username",
			expectSK:     "USER#example@org.com",
			expectEntity: "USER",
		},
		{
			name: "Happy - User without id",
			entity: func() itf.DBKeyer {
				u := entity.User{
					ID:         entity.IDFromStringOrNil(""),
					Email:      "example@org.com",
					Firstname:  "Linus",
					Lastname:   "Boehm",
					Attributes: nil,
				}
				return NewUserEntity(u)

			},
			expectPK:     "USER#",
			expectSK:     "USER#example@org.com",
			expectEntity: "USER",
		},
		{
			name: "Happy - Tenant without parent",
			entity: func() itf.DBKeyer {
				return NewTenantEntity(mainTenant)
			},
			expectPK:     fmt.Sprintf("TENANT#%s", mainTenant.ID.String()),
			expectSK:     fmt.Sprintf("TENANT#%s", mainTenant.ID.String()),
			expectEntity: "TENANT",
		},
		{
			name: "Happy - Tenant with parent",
			entity: func() itf.DBKeyer {
				return NewTenantEntity(subTenant)
			},
			expectPK:     fmt.Sprintf("TENANT#%s", mainTenant.ID.String()),
			expectSK:     fmt.Sprintf("TENANT#%s", subTenant.ID.String()),
			expectEntity: "TENANT",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := test.entity()

			assert.Equal(t, test.expectPK, e.GetPK().String())
			assert.Equal(t, test.expectSK, e.GetSK().String())
			assert.Equal(t, test.expectEntity, e.GetEntity().String())
		})
	}
}
