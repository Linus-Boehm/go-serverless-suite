package itf

import (
	"fmt"

	"github.com/Linus-Boehm/go-serverless-suite/entity"
)

//go:generate mockgen -destination=basetable_mocks.go -package=itf -source=basetable.go

type BaseTableProvider interface {
	PutItem(row DBKeyer) error
	RemoveItem(key DBKeyer, item DBKeyer) error
	RemoveItemSoft(key DBKeyer, item DeletableKey) error
	RemoveMainEntity(entity fmt.Stringer, id fmt.Stringer) error
	ReadItem(key DBKeyer, row DBKeyer) error
	ReadAllWithPK(key fmt.Stringer, index *entity.TableIndex, entity fmt.Stringer, rows interface{}) error
	ReadItemFromIndex(key DBKeyer, index *entity.TableIndex, row DBKeyer) error
	GetEntity(entityIndex entity.TableIndex, entity fmt.Stringer, rows interface{}, withDeleted bool) error
	DeleteTable() error
	TranslateDBError(err error, entity fmt.Stringer, id fmt.Stringer) error
	BatchReadItems(keys []DBKeyer, rows interface{}) error
	BatchWriteItems(rows ...interface{}) error
	WithIndex(index entity.TableIndex) BaseTableProvider
	WithDefaultIndices() BaseTableProvider
	BatchDeleteItems(rows []DBKeyer) (int, error)
}

type DBKeyer interface {
	GetPK() fmt.Stringer
	GetSK() fmt.Stringer
	GetEntity() fmt.Stringer
}

type DeletableKey interface {
	DBKeyer
	IsDeleted() bool
	SoftDeleteNow()
}
