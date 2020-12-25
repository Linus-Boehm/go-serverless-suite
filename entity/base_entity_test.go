package entity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntity(t *testing.T) {
	uid := NewEntityIDV4()

	mainTenant := NewTenant("Foo", "foo", ID{})
	subTenant := NewTenant("Bar", "bar", mainTenant.ID)

	tests := []struct {
		name         string
		entity       DBEntityGetter
		expectEntity string
		expectPK     string
		expectSK     string
	}{
		{
			name: "Happy - User with uuid",
			entity: &User{
				ID:         uid,
				Email:      "example@org.com",
				Firstname:  "Linus",
				Lastname:   "Boehm",
				Attributes: nil,
			},
			expectPK:     fmt.Sprintf("USER#%s", uid.String()),
			expectSK:     "USER#example@org.com",
			expectEntity: "USER",
		},
		{
			name: "Happy - User with string id",
			entity: &User{
				ID:         IDFromStringOrNil("username"),
				Email:      "example@org.com",
				Firstname:  "Linus",
				Lastname:   "Boehm",
				Attributes: nil,
			},
			expectPK:     "USER#username",
			expectSK:     "USER#example@org.com",
			expectEntity: "USER",
		},
		{
			name: "Happy - User without id",
			entity: &User{
				ID:         IDFromStringOrNil(""),
				Email:      "example@org.com",
				Firstname:  "Linus",
				Lastname:   "Boehm",
				Attributes: nil,
			},
			expectPK:     "USER#",
			expectSK:     "USER#example@org.com",
			expectEntity: "USER",
		},
		{
			name:         "Happy - Tenant without parent",
			entity:       mainTenant,
			expectPK:     fmt.Sprintf("TENANT#%s", mainTenant.ID.String()),
			expectSK:     fmt.Sprintf("TENANT#%s", mainTenant.ID.String()),
			expectEntity: "TENANT",
		},
		{
			name:         "Happy - Tenant with parent",
			entity:       subTenant,
			expectPK:     fmt.Sprintf("TENANT#%s", mainTenant.ID.String()),
			expectSK:     fmt.Sprintf("TENANT#%s", subTenant.ID.String()),
			expectEntity: "TENANT",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := test.entity.GetDBEntity()

			assert.Equal(t, test.expectPK, e.GetPK().String())
			assert.Equal(t, test.expectSK, e.GetSK().String())
			assert.Equal(t, test.expectEntity, e.GetEntity().String())
		})
	}
}
