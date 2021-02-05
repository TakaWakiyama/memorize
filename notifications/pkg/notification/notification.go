package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

//RequestBody is for sending notification on slack
type RequestBody struct {
	Text string `json:"text"`
}

//SendNotificationToSlack is
func SendNotificationToSlack(webhookURL string, text string) {
	if os.Getenv("LAMBDA_ENV_TYPE") != "prod" {
		return
	}
	body := RequestBody{}
	body.Text = text
	jsonBody, _ := json.Marshal(body)
	fmt.Println(jsonBody)
	res, _ := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonBody))
	fmt.Println(res)
}
