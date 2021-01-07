package persistence

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Linus-Boehm/go-serverless-suite/entity"

	db "github.com/Linus-Boehm/go-serverless-suite/infra/provider/db"

	"github.com/Linus-Boehm/go-serverless-suite/common"

	"github.com/Linus-Boehm/go-serverless-suite/itf"

	"github.com/guregu/dynamo"
)

var (
	DynamoMainIndex = entity.TableIndex{
		Name: "table",
		PK:   "pk",
		SK:   "sk",
	}
	DynamoReverseIndex = entity.TableIndex{
		Name: "gsi-1-reverse",
		PK:   "sk",
		SK:   "pk",
	}
	// DynamoEntityIndex is used for list queries based on an single GetEntity in a single-table layout eg. GetAllUsers
	DynamoEntityIndex = entity.TableIndex{
		Name: "gsi-2-entity",
		PK:   "entity",
		SK:   "slug",
	}
)

type dynamoBaseTable struct {
	DB        *db.DynamoProvider
	Table     dynamo.Table
	mainIndex entity.TableIndex
	indexes   map[string]entity.TableIndex
}

func NewTestProvider(from interface{}) (itf.BaseTableProvider, error) {
	dbProvider, err := db.ConnectDynamo(&db.DynamoConfig{
		Endpoint:   common.StringPtr("http://localhost:8000"),
		Region:     common.StringPtr("eu-central-1"),
		AwsProfile: "",
	})
	if err != nil {
		panic(err)
	}
	id := entity.NewEntityIDV4()
	tbl := dbProvider.TableFromName("go-svl", "basetable", fmt.Sprintf("test-%s", id.String()))
	err = tbl.DeleteTable().Run()
	if err != nil {
		fmt.Println(err)
	}
	err = dbProvider.DB().
		CreateTable(tbl.Name(), from).
		Provision(4, 4).
		Run()
	return NewFromDynamoDB(dbProvider, tbl, nil).WithDefaultIndices(), err
}

func NewFromDynamoDB(provider *db.DynamoProvider, table dynamo.Table, mainIndex *entity.TableIndex) itf.BaseTableProvider {
	b := &dynamoBaseTable{
		DB:      provider,
		Table:   table,
		indexes: map[string]entity.TableIndex{},
	}
	if mainIndex == nil {
		mainIndex = &DynamoMainIndex
	}
	b.indexes["main"] = *mainIndex
	b.mainIndex = *mainIndex
	return b
}

func (b *dynamoBaseTable) WithIndex(index entity.TableIndex) itf.BaseTableProvider {
	b.indexes[index.Name] = index
	return b
}

func (b *dynamoBaseTable) WithDefaultIndices() itf.BaseTableProvider {
	b.indexes[DynamoReverseIndex.Name] = DynamoReverseIndex
	b.indexes[DynamoEntityIndex.Name] = DynamoEntityIndex
	return b
}

func (b *dynamoBaseTable) PutItem(row itf.DBKeyer) error {
	err := b.Table.Put(row).Run()
	return b.TranslateDBError(err, row.GetEntity(), row.GetPK())
}

func (b *dynamoBaseTable) RemoveMainEntity(entity fmt.Stringer, id fmt.Stringer) error {
	key := JoinKeys(entity.String(), id.String())
	err := b.Table.Delete(b.mainIndex.PK, key).Range(b.mainIndex.SK, key).Run()
	return b.TranslateDBError(err, entity, id)
}

func (b *dynamoBaseTable) RemoveItem(key itf.DBKeyer, item itf.DBKeyer) error {
	err := b.Table.Delete(b.mainIndex.PK, key.GetPK().String()).Range(b.mainIndex.SK, key.GetSK().String()).OldValue(item)
	return b.TranslateDBError(err, key.GetEntity(), key.GetPK())
}

func (b *dynamoBaseTable) RemoveItemSoft(key itf.DBKeyer, item itf.DeletableKey) error {
	err := b.ReadItem(key, item)
	if err != nil {
		return err
	}
	item.SoftDeleteNow()
	return b.PutItem(item)
}

func (b *dynamoBaseTable) RemoveAllWithPK(entity fmt.Stringer, id fmt.Stringer) error {
	key := JoinKeys(entity.String(), id.String())
	err := b.Table.Delete(b.mainIndex.PK, key).Run()
	return b.TranslateDBError(err, entity, id)
}

func (b *dynamoBaseTable) ReadItem(key itf.DBKeyer, item itf.DBKeyer) error {
	return b.ReadItemFromIndex(key, nil, item)
}

func (b *dynamoBaseTable) ReadAllWithPK(key fmt.Stringer, index *entity.TableIndex, entity fmt.Stringer, rows interface{}) error {
	if index == nil {
		index = &b.mainIndex
	}
	query := b.Table.Get(index.PK, key.String())
	if index.Name != b.mainIndex.Name {
		query = query.Index(index.Name)
	}
	if entity.String() != "" {
		query = query.Filter("'entity' = ?", entity.String())
	}

	err := query.All(rows)
	if err != nil {
		return b.TranslateDBError(err, nil, key)
	}
	return nil
}

func (b *dynamoBaseTable) ReadItemFromIndex(key itf.DBKeyer, index *entity.TableIndex, row itf.DBKeyer) error {
	if index == nil {
		index = &b.mainIndex
	}
	query := b.Table.Get(index.PK, key.GetPK().String()).Range(index.SK, dynamo.Equal, key.GetSK().String())
	if index.Name != b.mainIndex.Name {
		query = query.Index(index.Name)
	}
	err := query.One(row)

	if err != nil {
		return b.TranslateDBError(err, key.GetEntity(), key.GetPK())
	}
	return nil
}

func (b *dynamoBaseTable) GetEntity(entityIndex entity.TableIndex, entity fmt.Stringer, rows interface{}, withDeleted bool) error {
	query := b.Table.Get(entityIndex.PK, entity.String()).Index(entityIndex.Name)
	if !withDeleted {
		query = query.Filter("attribute_not_exists(deleted_at) OR 'deleted_at' = ?", 0)
	}
	err := query.All(rows)
	if err != nil {
		return b.TranslateDBError(err, entity, nil)
	}
	return nil
}

func (b *dynamoBaseTable) DeleteTable() error {
	return b.Table.DeleteTable().Run()
}

func (b *dynamoBaseTable) TranslateDBError(err error, entity fmt.Stringer, id fmt.Stringer) error {
	if err == nil {
		return nil
	}
	if errors.As(err, &dynamo.ErrNotFound) {
		return common.NewEntityNotFoundError(id, entity)
	}
	return err
}

func JoinKeys(fragments ...string) string {
	return strings.Join(fragments, "#")
}
