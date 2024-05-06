package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.DynamoDBEvent) (string, error) {

	var res string
	for _, record := range event.Records {
		log.Printf("Stream record: %v", record)

		if record.EventName == "INSERT" {
			who := record.Change.NewImage["Username"].String()
			when := record.Change.NewImage["Timestamp"].String()
			what := record.Change.NewImage["Message"].String()
			log.Printf("New record: %s %s %s", who, when, what)
			// who when whatを結合してresに格納
			res = when + ", " + who + ", " + what
		}
	}

	return res, nil
}

func main() {
	lambda.Start(handler)
}
