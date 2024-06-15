// https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html
package main

import (
	"context"
	"fmt"
	"log"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tkhrk1010/infra-samples/lambda-container/normal/lambda/internal/model"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event *MyEvent) (*string, error) {
	if event == nil {
		log.Println("received nil event")
		return nil, fmt.Errorf("received nil event")
	}
	user := model.NewUser(event.Name)
	message := fmt.Sprintf("Hello %s!", user.Name)
	log.Printf("message: %s", message)
	return &message, nil
}

func main() {
	lambda.Start(HandleRequest)
}
