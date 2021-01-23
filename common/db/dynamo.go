package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"os"
)

// InitalizeDynamoClient is initalze Dynamo
func InitalizeDynamoClient() *dynamo.DB {
	region := "ap-northeast-1"
	var config aws.Config
	config.Region = aws.String(region)
	if os.Getenv("ENV_TYPE") != "production" {
		config.Endpoint = aws.String("http://127.0.0.1:8000")
		config.Credentials = credentials.NewStaticCredentials("dummy", "dummy", "dummy")
	}

	db := dynamo.New(session.New(), &config)
	return db
}
