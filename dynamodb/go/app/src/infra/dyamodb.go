package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// DynamoDBクライアントの初期化
func InitializeDynamoDBClient(ctx context.Context) *dynamodb.Client {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == dynamodb.ServiceID {
			return aws.Endpoint{
				PartitionID:   "aws",
				// localStackが異なるdocker network上にあるため、host.docker.internalを使ってDockerからMacのlocalhostを経由してアクセスする
				URL:           "http://host.docker.internal:4566",
				SigningRegion: "us-east-1",
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(aws.AnonymousCredentials{}),
		config.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		panic(fmt.Sprintf("unable to load SDK config, %v", err))
	}

	return dynamodb.NewFromConfig(cfg)
}

// CreateTable - テーブルの作成
func CreateTable(svc *dynamodb.Client, tableName string) {
	createTableInput := &dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       types.KeyTypeHash,
			},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("ID"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	_, err := svc.CreateTable(context.TODO(), createTableInput)
	if err != nil {
		panic("failed to create table, " + err.Error())
	}
	fmt.Println("Table created:", tableName)
}

// PutItem - アイテムの追加
func PutItem(svc *dynamodb.Client, tableName string, item map[string]types.AttributeValue) {
	putItemInput := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	}

	_, err := svc.PutItem(context.TODO(), putItemInput)
	if err != nil {
		panic("failed to put item, " + err.Error())
	}
	fmt.Println("Item added to table:", tableName)
}

// GetItem - アイテムの取得
func GetItem(svc *dynamodb.Client, tableName string, key map[string]types.AttributeValue) {
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	}

	res, err := svc.GetItem(context.TODO(), getItemInput)
	if err != nil {
		panic("failed to get item, " + err.Error())
	}

	itemBytes, err := json.Marshal(res.Item)
	if err != nil {
		panic("failed to marshal item, " + err.Error())
	}
	fmt.Printf("Item: %s\n", itemBytes)
}
