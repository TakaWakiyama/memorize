package main

import (
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/guregu/dynamo"

	"memos/common/db"
	"memos/settings/pkg/handlers"
)

var (
	err        error
	dynaClient dynamo.DB
)

func main() {
	dynaClient = *db.InitalizeDynamoClient()
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case "POST":
		return handlers.CreateSetting(request, &dynaClient)
	default:
		return events.APIGatewayProxyResponse{}, errors.New("Method is not allowed")
	}
}
