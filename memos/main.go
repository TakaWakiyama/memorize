package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

	"memos/common/db"
)

var (
	err        error
	dynaClient dynamodbiface.DynamoDBAPI
)

func main() {
	dynaClient, err = db.InitalizeDynamoClient()
	if err != nil {
		return
	}
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	/*
		switch req.HTTPMethod {
		case "GET":
			// return handlers.GetUser(req, tableName, dynaClient)
		case "POST":
			// return handlers.CreateUser(req, tableName, dynaClient)
		case "PUT":
			// return handlers.UpdateUser(req, tableName, dynaClient)
		case "DELETE":
			// return handlers.DeleteUser(req, tableName, dynaClient)
		default:
			// return handlers.UnhandledMethod()
		}
	*/
	fmt.Print(request.PathParameters)
	fmt.Print(request.QueryStringParameters)
	res := events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", request.PathParameters["pk"]),
		StatusCode: 200,
	}

	return res, nil
}
