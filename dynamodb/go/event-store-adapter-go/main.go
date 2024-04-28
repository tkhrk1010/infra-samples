package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	esag "github.com/j5ik2o/event-store-adapter-go/pkg"

	"github.com/tkhrk1010/infra-samples/dynamodb/go/event-store-adapter-go/domain"
	"github.com/tkhrk1010/infra-samples/dynamodb/go/event-store-adapter-go/domain/models"
	"github.com/tkhrk1010/infra-samples/dynamodb/go/event-store-adapter-go/interfaceAdaptor/repository"
)

func initLogger() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func newDynamoDBClient() *dynamodb.Client {
	//
	// dynamoDB client
	awsDynamoDBEndpointUrl := "http://localhost:4566"
	// localstackはus-east-1のみ対応
	awsRegion := "us-east-1"
	awsDynamoDBAccessKeyId := "dummy"
	awsDynamoDBSecretKey := "dummy"
	customResolver := aws.EndpointResolverWithOptionsFunc(
		func(service, region string, opts ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsDynamoDBEndpointUrl,
				SigningRegion: region,
			}, nil
		})
	awsCfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(awsRegion),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsDynamoDBAccessKeyId, awsDynamoDBSecretKey, "x")),
	)
	if err != nil {
		panic(err)
	}

	return dynamodb.NewFromConfig(awsCfg)
}

func newUserAccountRepository(dynamodbClient *dynamodb.Client) *repository.UserAccountRepository {
	journalTableName := "journal"
	snapshotTableName := "snapshot"
	journalAidIndexName := "journal-aid-index"
	snapshotAidIndexName := "snapshot-aid-index"
	shardCount := 1

	eventStore, err := esag.NewEventStoreOnDynamoDB(
		dynamodbClient,
		journalTableName,
		snapshotTableName,
		journalAidIndexName,
		snapshotAidIndexName,
		uint64(shardCount),
		repository.EventConverter,
		repository.SnapshotConverter,
		esag.WithEventSerializer(repository.NewEventSerializer()),
		esag.WithSnapshotSerializer(repository.NewSnapshotSerializer()))
	if err != nil {
		panic(err)
	}
	return repository.NewUserAccountRepository(eventStore)
}

func main() {
	fmt.Println("start")

	initLogger()

	dynamodbClient := newDynamoDBClient()

	repository := newUserAccountRepository(dynamodbClient)

	// jurnalとsnapshot tableのrecordを全部消すようにしたい
	// localで試す用
	// err = eventStore.ClearAll()

	fmt.Println("NewUserAccount")
	userAccount1, userAccountCreated := domain.NewUserAccount("test", models.NewUserAccountId())
	fmt.Printf("userAccount1 = %+v\n", userAccount1)

	// Store an aggregate with a create event
	fmt.Println("StoreEventAndSnapshot userAccountCreated")
	err := repository.StoreEventAndSnapshot(userAccountCreated, &userAccount1)
	if err != nil {
		panic(err)
	}

	// Replay the aggregate from the event store
	fmt.Println("FindById userAccount1.id")
	userAccount2, err := repository.FindById(userAccount1.GetId())
	if err != nil {
		panic(err)
	}
	fmt.Printf("userAccount2 = %+v\n", userAccount2)

	// Execute a command on the aggregate
	fmt.Println("ChangeName")
	userAccountUpdated, userAccountNameChanged := userAccount2.ChangeName("test2")
	fmt.Printf("userAccountUpdated = %+v\n", userAccountUpdated)

	// Store the new event without a snapshot
	fmt.Println("StoreEvent userAccountNameChanged")
	// err = repository.StoreEvent(userAccountNameChanged, userAccountUpdated.GetVersion())
	// Store the new event with a snapshot
	err = repository.StoreEventAndSnapshot(userAccountNameChanged, userAccountUpdated)
	if err != nil {
		panic(err)
	}

	fmt.Println("end")
}
