package main

import (
	"errors"
	"fmt"
	"time"

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

func getCurrentDate() string {
	t := time.Now()
	const layout2 = "2006-01-02"
	fmt.Println(t.Format(layout2))
	return t.Format(layout2)
}

func main() {
	fmt.Println(getCurrentDate())

	dynaClient = *db.InitalizeDynamoClient()
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	/*
		case "GET":
			return handlers.GetMemo(&dynaClient, "gggggg")
	*/
	case "POST":
		return handlers.CreateMemo(&dynaClient, request)
	default:
		return nil, errors.New("ggg")
	}
}
