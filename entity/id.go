package entity

import (
	"fmt"
	"strings"

	"github.com/Linus-Boehm/go-serverless-suite/common"
	"github.com/google/uuid"
)

// ID is a wrapper around google/uuid to allow a seamlessly work between the API and persistence layer
// It allows the use of UUIDs and from normal string IDs this makes the use of it very flexible.
type ID struct {
	id  *uuid.UUID
	val *string
}

func MustEntityID(id *ID, err error) ID {
	if id == nil || err != nil {
		panic("EntityID must be set")
	}
	return *id
}

func NewEntityIDV4() ID {
	i := uuid.Must(uuid.NewRandom())
	return ID{id: &i}
}

func IDFromStringOrNil(id string) ID {
	if id == "" {
		return ID{}
	}
	uid, err := uuid.Parse(id)
	if err == nil && uuid.Nil.String() != uid.String() {
		return ID{
			id: &uid,
		}
	}
	return ID{val: &id}
}

// IDFromDBString converts string in the form of `GetEntity#0000` (where 0000 is a UUID or a other string)
// to a ID object. It return an error if it could not split entityStr.
func IDFromDBString(entityStr string) (ID, error) {
	parts := strings.Split(entityStr, "#")
	if len(parts) < 2 {
		return ID{}, fmt.Errorf("not a valid EntityID! missing GetEntity part separated with # in %s", entityStr)
	}
	return IDFromStringOrNil(parts[1]), nil
}

func (e ID) IsNil() bool {
	if e.val != nil {
		return *e.val == ""
	}
	if e.id != nil {
		return e.id.String() == uuid.Nil.String() || *e.id == uuid.Nil
	}
	return true
}

func (e ID) String() string {
	if e.id == nil {
		return common.StringValue(e.val)
	}
	return e.id.String()
}

func (e ID) NewV4() {
	i := uuid.Must(uuid.NewRandom())
	e.id = &i
}

func (e ID) NewV4IfEmpty() {
	if e.id == nil && e.val == nil {
		e.NewV4()
	}
}

func (e ID) WithEntity(entity EntityKey) fmt.Stringer {
	return common.NewJoinedStringDBKey(entity.String(), e.String())
}
