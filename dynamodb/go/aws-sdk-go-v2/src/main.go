package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/tkhrk1010/infra-samples/dynamodb/go/aws-sdk-go-v2/src/infra"
)

func main() {
	ctx := context.TODO()

	// DynamoDBクライアントの初期化
	svc := infra.InitializeDynamoDBClient(ctx)

	// テーブル名の指定
	tableName := "MyTable"

	// テーブルの作成
	infra.CreateTable(svc, tableName)

	// アイテムの追加
	item := map[string]types.AttributeValue{
		"ID":   &types.AttributeValueMemberS{Value: "1"},
		"Name": &types.AttributeValueMemberS{Value: "John Doe"},
	}
	infra.PutItem(svc, tableName, item)

	// アイテムの取得
	key := map[string]types.AttributeValue{
		"ID": &types.AttributeValueMemberS{Value: "1"},
	}
	infra.GetItem(svc, tableName, key)
}
