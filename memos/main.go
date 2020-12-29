package main

import (
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/guregu/dynamo"

	"memos/common/db"
	"memos/memos/pkg/handlers"
)

var (
	err        error
	dynaClient dynamo.DB
)

func main() {
	dynaClient = *db.InitalizeDynamoClient()
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Printf("called")
	switch request.HTTPMethod {
	case "GET":
		return handlers.GetMemo(&dynaClient, "gggggg")
	case "POST":
		return handlers.CreateMemo(&dynaClient)
	default:
		return nil, errors.New("ggg")
	}
}
