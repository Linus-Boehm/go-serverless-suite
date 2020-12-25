package itf

import (
	"fmt"

	"github.com/Linus-Boehm/go-serverless-suite/entity"
)

//go:generate mockgen -destination=basetable_mocks.go -package=itf -source=basetable.go

type BaseTableProvider interface {
	PutItem(row entity.DBProvider) error
	RemoveItem(key entity.DBKeyer) error
	RemoveMainEntity(entity fmt.Stringer, id fmt.Stringer) error
	ReadItem(key entity.DBKeyer, row entity.DBProvider) error
	ReadAllWithPK(key fmt.Stringer, index *entity.TableIndex, entity fmt.Stringer, rows interface{}) error
	ReadItemFromIndex(key entity.DBKeyer, index *entity.TableIndex, row entity.DBProvider) error
	GetEntity(entityIndex entity.TableIndex, entity fmt.Stringer, rows interface{}, withDeleted bool) error
	DeleteTable() error
	TranslateDBError(err error, entity fmt.Stringer, id fmt.Stringer) error
	BatchReadItems(keys []entity.DBKeyer, rows interface{}) error
	BatchWriteItems(rows ...interface{}) error
	WithIndex(index entity.TableIndex) BaseTableProvider
	BatchDeleteItems(rows []entity.DBKeyer) (int, error)
}
