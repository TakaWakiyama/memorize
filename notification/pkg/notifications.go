package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//RequestBody is for sending notification on slack
type RequestBody struct {
	Text string `json:"text"`
}

//SendNotificationToSlack is
func SendNotificationToSlack(webhookURL string, text string) {
	body := RequestBody{}
	body.Text = text
	jsonBody, _ := json.Marshal(body)

	res, _ := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonBody))
	fmt.Println(res)

}
