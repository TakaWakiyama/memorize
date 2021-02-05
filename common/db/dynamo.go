package db

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// InitalizeDynamoClient is initalze Dynamo
func InitalizeDynamoClient() *dynamo.DB {
	region := "ap-northeast-1"
	var config aws.Config
	config.Region = aws.String(region)
	if os.Getenv("LAMBDA_ENV_TYPE") == "local" {
		config.Endpoint = aws.String("http://172.22.0.1:8000")
		config.Credentials = credentials.NewStaticCredentials("dummy", "dummy", "dummy")
	}

	db := dynamo.New(session.New(), &config)
	return db
}
