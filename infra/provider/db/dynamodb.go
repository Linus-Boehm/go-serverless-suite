package db

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type DynamoConfig struct {
	Endpoint   *string
	Region     *string
	AwsProfile string
}

type DynamoProvider struct {
	db     *dynamo.DB
	config *DynamoConfig
}

type DynamoTableIndex struct {
	Name string
	Pk   string
	Sk   string
}

func ConnectDynamo(config *DynamoConfig) (*DynamoProvider, error) {
	conf := aws.Config{
		Endpoint: config.Endpoint,
		Region:   config.Region,
	}
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: config.AwsProfile,
		Config:  conf,
	})
	if err != nil {
		return nil, err
	}
	db := dynamo.New(sess, &conf)

	return &DynamoProvider{
		db:     db,
		config: config,
	}, nil
}

func (d *DynamoProvider) Table(name string) dynamo.Table {
	return d.db.Table(name)
}

func (d *DynamoProvider) TableFromName(app string, tablename string, stage string) dynamo.Table {
	name := fmt.Sprintf("%s-%s-%s", app, tablename, stage)
	return d.Table(name)
}

func (d *DynamoProvider) DB() *dynamo.DB {
	return d.db
}
