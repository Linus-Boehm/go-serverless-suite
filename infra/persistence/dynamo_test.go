package persistence

import (
	"errors"
	"fmt"
	"testing"

	"github.com/guregu/dynamo"

	"github.com/Linus-Boehm/go-serverless-suite/common"

	"github.com/Linus-Boehm/go-serverless-suite/entity"
	"github.com/stretchr/testify/assert"
)

type TestEntity struct {
	PK         string            `dynamo:"pk,hash" index:"gsi-1-reverse,range"`
	SK         string            `dynamo:"sk,range" index:"gsi-1-reverse,hash"`
	Entity     string            `dynamo:"entity,omitempty" index:"gsi-2-entity,hash"`
	Slug       string            `dynamo:"slug,omitempty" index:"gsi-2-entity,range"`
	Timestamps entity.Timestamps `dynamo:"timestamps"`
}

func (e *TestEntity) GetPK() fmt.Stringer {
	return common.NewString(e.PK)
}

func (e *TestEntity) GetSK() fmt.Stringer {
	return common.NewString(e.SK)
}

func (e *TestEntity) GetEntity() fmt.Stringer {
	return common.NewString("TEST")
}

func TestProvider(t *testing.T) {
	b, err := NewTestProvider(TestEntity{})
	assert.NoError(t, err)
	assert.NoError(t, b.DeleteTable())
}

func TestDynamoBaseTable_PutReadItem(t *testing.T) {
	b, err := NewTestProvider(TestEntity{})
	assert.NoError(t, err)
	defer b.DeleteTable()
	ts := entity.Timestamps{}
	ts.CreatedNow()
	e := &TestEntity{
		PK:         entity.NewEntityIDV4().String(),
		SK:         "example.org",
		Entity:     "TEST",
		Slug:       "bla",
		Timestamps: ts,
	}
	err = b.PutItem(e)
	assert.NoError(t, err)
	oldU := e
	newU := TestEntity{}
	err = b.ReadItem(e, &newU)
	assert.NoError(t, err)

	assert.Equal(t, oldU.PK, newU.PK)
	assert.Equal(t, oldU.SK, newU.SK)
}

func TestDynamoBaseTable_ReadAllWithPK(t *testing.T) {
	b, err := NewTestProvider(TestEntity{})
	assert.NoError(t, err)
	defer b.DeleteTable()
	ts := entity.Timestamps{}
	ts.CreatedNow()
	e := &TestEntity{
		PK:         entity.NewEntityIDV4().String(),
		SK:         "example.org",
		Entity:     "TEST",
		Slug:       "bla",
		Timestamps: ts,
	}
	err = b.PutItem(e)
	assert.NoError(t, err)
	result := []TestEntity{}
	err = b.ReadAllWithPK(e.GetPK(), nil, common.NewString("TEST"), &result)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, result[0].PK, e.GetPK().String())
}

func TestDynamoBaseTable_BatchWriteItems(t *testing.T) {
	b, err := NewTestProvider(TestEntity{})
	assert.NoError(t, err)
	defer b.DeleteTable()
	newU := func() TestEntity {
		return TestEntity{
			PK:     entity.NewEntityIDV4().String(),
			SK:     "example.org",
			Entity: "TEST",
			Slug:   "bla",
		}
	}
	var payload []interface{}

	for i := 0; i < 50; i++ {
		u := newU()
		payload = append(payload, &u)
	}

	err = b.BatchWriteItems(payload...)
	assert.NoError(t, err)
	var counter []TestEntity
	err = b.GetEntity(DynamoEntityIndex, (&TestEntity{}).GetEntity(), &counter, false)
	assert.NoError(t, err)
	assert.Len(t, counter, 50)
}

func Test_dynamoBaseTable_TranslateDBError(t *testing.T) {
	type args struct {
		entity   fmt.Stringer
		id       fmt.Stringer
		inputErr error
	}
	tests := []struct {
		name         string
		args         args
		wantErrEqual error
		wantErr      bool
	}{
		{
			name: "happy return domain err",
			args: args{
				entity:   common.NewString("TEST"),
				id:       common.NewString("ID"),
				inputErr: dynamo.ErrNotFound,
			},

			wantErrEqual: common.NewEntityNotFoundError(common.NewString("ID"), common.NewString("TEST")),
			wantErr:      true,
		},
		{
			name: "happy return other errors",
			args: args{
				entity:   common.NewString("TEST"),
				id:       common.NewString("ID"),
				inputErr: errors.New("foo"),
			},

			wantErrEqual: errors.New("foo"),
			wantErr:      true,
		},
		{
			name: "happy return other errors",
			args: args{
				entity:   common.NewString("TEST"),
				id:       common.NewString("ID"),
				inputErr: nil,
			},

			wantErrEqual: nil,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := NewTestProvider(TestEntity{})
			assert.NoError(t, err)
			defer b.DeleteTable()

			gotErr := b.TranslateDBError(tt.args.inputErr, tt.args.entity, tt.args.id)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("TranslateDBError() error = %v, wantErr %v", gotErr, tt.wantErr)
			}
			if tt.wantErr == true {
				if !errors.As(gotErr, &tt.wantErrEqual) {
					t.Errorf("TranslateDBError() gotError = %v, want %v", gotErr, tt.wantErrEqual)
				}
			}
		})
	}
}
