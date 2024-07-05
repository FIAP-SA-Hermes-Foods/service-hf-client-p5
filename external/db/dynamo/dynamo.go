package dynamo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type dynamoDB struct {
	ctx    context.Context
	config aws.Config
	client *dynamodb.Client
}

func NewDynamoDB(config aws.Config) *dynamoDB {
	return &dynamoDB{config: config}
}

func (d *dynamoDB) voucherDynamo() {
	d.client = dynamodb.NewFromConfig(d.config)
}

func (d *dynamoDB) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	if d.client == nil {
		d.voucherDynamo()
	}
	if d.ctx == nil {
		d.ctx = context.Background()
	}
	return d.client.Scan(d.ctx, input)
}

func (d *dynamoDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	if d.client == nil {
		d.voucherDynamo()
	}
	if d.ctx == nil {
		d.ctx = context.Background()
	}
	return d.client.PutItem(d.ctx, input)
}

func (d *dynamoDB) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	if d.client == nil {
		d.voucherDynamo()
	}
	if d.ctx == nil {
		d.ctx = context.Background()
	}
	return d.client.Query(d.ctx, input)
}
