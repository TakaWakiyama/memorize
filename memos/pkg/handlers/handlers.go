package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/guregu/dynamo"

	"memos/memos/pkg/memos"
)

const tableName = "Memos"

var (
	err error
)

// GetMemo is get memo by pk
func GetMemo(dynaClient *dynamo.DB, pk string) (*events.APIGatewayProxyResponse, error) {
	// dynamo.Table
	table := dynaClient.Table(tableName)
	memo, err := memos.Get(table, pk)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	body := map[string]string{
		"memo_id":   memo.MemoID,
		"memo_type": memo.MemoType,
		"value":     memo.Value,
	}
	resp, err := apires(body)
	return resp, err
}

// CreateMemo is create a memo
func CreateMemo(dynaClient *dynamo.DB) (*events.APIGatewayProxyResponse, error) {
	table := dynaClient.Table(tableName)
	memo := memos.Memo{
		MemoID:   "gggggg",
		MemoType: "Golang",
		Value:    "I think it's so difficult",
	}
	memoID, err := memos.Create(table, &memo)
	if err != nil {
		return nil, err
	}
	body := map[string]string{"memo_id": memoID}
	resp, err := apires(body)
	return resp, err
}

func apires(body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = 200

	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	return &resp, nil
}
