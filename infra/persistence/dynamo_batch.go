package persistence

import (
	"github.com/Linus-Boehm/go-serverless-suite/entity"
	"github.com/guregu/dynamo"
)

func (b *dynamoBaseTable) BatchWriteItems(rows ...interface{}) error {
	_, err := b.Table.Batch(b.mainIndex.PK, b.mainIndex.SK).Write().Put(rows...).Run()
	if err != nil {
		return err
	}
	return nil
}

func (b *dynamoBaseTable) BatchDeleteItems(rows []entity.DBKeyer) (int, error) {
	keys := make([]dynamo.Keyed, len(rows))
	for i, row := range rows {
		keys[i] = dynamo.Keys{row.GetPK().String(), row.GetSK().String()}
	}
	num, err := b.Table.Batch(b.mainIndex.PK, b.mainIndex.SK).Write().Delete(keys...).Run()
	return num, err
}

func (b *dynamoBaseTable) BatchReadItems(keys []entity.DBKeyer, rows interface{}) error {
	//TODO split after 25
	keyed := make([]dynamo.Keyed, len(keys))
	for i, key := range keys {
		keyed[i] = dynamo.Keys{key.GetPK().String(), key.GetSK().String()}
	}
	err := b.Table.Batch(b.mainIndex.PK, b.mainIndex.SK).Get(keyed...).All(rows)
	if err != nil {
		return b.TranslateDBError(err, nil, nil)
	}
	return nil
}
