package handlers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/guregu/dynamo"

	"memos/memos/pkg/memos"
)

const tableName = "Memos"

var (
	err error
)

/*
func GetMemo(dynaClient *dynamo.DB, pk string) (*events.APIGatewayProxyResponse, error) {
	// dynamo.Table
	table := dynaClient.Table(tableName)
	memo, err := memos.Get(table, pk)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	body := map[string]string{
		"memo_id":   memo.MemoId,
		"memo_type": memo.MemoType,
	}
	resp, err := apires(body)
	return resp, err
}
*/

// CreateMemo is create a memo
func CreateMemo(dynaClient *dynamo.DB, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var memo memos.Memo
	table := dynaClient.Table(tableName)
	json.Unmarshal([]byte(request.Body), &memo)
	memo.User = "Twaki" // next -> get auth info from User
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
