package persistence

import (
	"github.com/Linus-Boehm/go-serverless-suite/itf"
	"github.com/guregu/dynamo"
)

func (b *dynamoBaseTable) BatchWriteItems(rows ...interface{}) error {
	_, err := b.Table.Batch(b.mainIndex.PK, b.mainIndex.SK).Write().Put(rows...).Run()
	if err != nil {
		return b.TranslateDBError(err, nil, nil)
	}
	return nil
}

func (b *dynamoBaseTable) BatchDeleteItems(rows []itf.DBKeyer) (int, error) {
	keys := make([]dynamo.Keyed, len(rows))
	for i, row := range rows {
		keys[i] = dynamo.Keys{row.GetPK().String(), row.GetSK().String()}
	}
	num, err := b.Table.Batch(b.mainIndex.PK, b.mainIndex.SK).Write().Delete(keys...).Run()
	return num, b.TranslateDBError(err, nil, nil)
}

func (b *dynamoBaseTable) BatchReadItems(keys []itf.DBKeyer, rows interface{}) error {
	//TODO split after 25
	uniqueKeys := UniqueKeys(keys)
	keyed := make([]dynamo.Keyed, len(uniqueKeys))
	for i, key := range uniqueKeys {
		keyed[i] = dynamo.Keys{key.GetPK().String(), key.GetSK().String()}
	}
	err := b.Table.Batch(b.mainIndex.PK, b.mainIndex.SK).Get(keyed...).All(rows)
	if err != nil {
		return b.TranslateDBError(err, nil, nil)
	}
	return nil
}

func UniqueKeys(keys []itf.DBKeyer) []itf.DBKeyer {
	unique := make(map[string]itf.DBKeyer)
	for _, key := range keys {
		unique[key.GetPK().String()+key.GetSK().String()] = key
	}
	result := make([]itf.DBKeyer, len(unique))
	i := 0
	for _, key := range unique {
		result[i] = key
		i++
	}
	return result
}
