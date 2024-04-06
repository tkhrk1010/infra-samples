package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func main() {
	ctx := context.TODO()

	// Custom Resolver for LocalStack DynamoDB
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == dynamodb.ServiceID {
			return aws.Endpoint{
				PartitionID: "aws",
				// localStackが異なるdocker network上にあるため、host.docker.internalを使ってDockerからMacのlocalhostを経由してアクセスする
				URL: "http://host.docker.internal:4566", // LocalStackのエンドポイント
				SigningRegion: "us-east-1",
			}, nil
		}
		// Fallback to default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	// 適当な認証情報を設定
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		// dev環境の場合は、AnonymousCredentialsを使うと楽。
		// ~/.aws/credentialsを使う場合は、config.LoadDefaultConfig()のみでOK
		config.WithCredentialsProvider(aws.AnonymousCredentials{}), // Anonymous credentials
		config.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		panic(fmt.Sprintf("unable to load SDK config, %v", err))
	}

	// DynamoDBクライアントの作成
	svc := dynamodb.NewFromConfig(cfg)

	// テーブルの作成
	createTableInput := &dynamodb.CreateTableInput{
		TableName: aws.String("MyTable"),
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

	_, err = svc.CreateTable(ctx, createTableInput)
	if err != nil {
		panic("failed to create table, " + err.Error())
	}
	fmt.Println("Table created")

	// アイテムの追加（Put）
	putItemInput := &dynamodb.PutItemInput{
		TableName: aws.String("MyTable"),
		Item: map[string]types.AttributeValue{
			"ID":   &types.AttributeValueMemberS{Value: "1"},
			"Name": &types.AttributeValueMemberS{Value: "John Doe"},
		},
	}

	_, err = svc.PutItem(ctx, putItemInput)
	if err != nil {
		panic("failed to put item, " + err.Error())
	}
	fmt.Println("Item added")

	// アイテムの取得（Get）
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String("MyTable"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: "1"},
		},
	}

	res, err := svc.GetItem(ctx, getItemInput)
	if err != nil {
		panic("failed to get item, " + err.Error())
	}
	
	// 取得したアイテムをJSON形式の文字列に変換して出力
	itemBytes, err := json.Marshal(res.Item)
	if err != nil {
		panic("failed to marshal item, " + err.Error())
	}
	fmt.Printf("Item: %s\n", itemBytes)
}
