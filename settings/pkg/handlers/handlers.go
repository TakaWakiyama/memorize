package handlers

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/guregu/dynamo"

	"memos/settings/pkg/settings"
)

const tableName = "Settings"

var (
	err error
)

func CreateSetting(req events.APIGatewayProxyRequest, dynaClient *dynamo.DB) (events.APIGatewayProxyResponse, error) {
	type Body struct {
		User              string `json:"user"`
		Category          string `json:"category"`
		ExecutionInterval []int  `json:"execution_interval"`
		TimeExecute       string `json:"time_execute"`
		Template          string `json:"template"`
	}
	jsonBytes := []byte(req.Body)
	// byte[]body
	body := new(Body)
	if err := json.Unmarshal(jsonBytes, body); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	setting := settings.Setting{
		User:              body.User,
		Category:          body.Category,
		ExecutionInterval: body.ExecutionInterval,
		TimeExecute:       body.TimeExecute,
		Template:          body.Template,
		IsActive:          1,
	}

	if !(setting.User != "" && setting.Category != "") {
		return events.APIGatewayProxyResponse{}, errors.New("ggg")
	}
	// user の存在と categoryの存在をチェック
	table := dynaClient.Table(tableName)
	settingKey, err := settings.Create(table, setting)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = 200

	stringBody, _ := json.Marshal(settingKey)
	resp.Body = string(stringBody)
	return resp, nil
}
