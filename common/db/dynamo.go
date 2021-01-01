package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// InitalizeDynamoClient is initalze Dynamo
func InitalizeDynamoClient() *dynamo.DB {
	region := "ap-northeast-1"
	db := dynamo.New(session.New(), &aws.Config{Region: aws.String(region)})

	return db
}
