package main

// event から ユーザー名, テーブル名を取得 , 経過日数の配列, テンプレートの取得
// 経過日数(相対)を日付に直す
// 日付毎に batch get で取得したアイテム (最大50 25件)を　テンプレートをもとにテキストシリアライズ
// slack の webhook urlにpost リクエストを送信する
import (
	"context"
	"log"
	"memos/common/db"
	"os"

	"memos/notifications/pkg/handlers"
	"memos/notifications/pkg/notification"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/guregu/dynamo"
)

var (
	dynaClient dynamo.DB
	err        error
)

// MyEvent is passed from CluodWatch
type MyEvent struct {
	User            string   `json:"user"`
	ItemType        string   `json:"item_type"`
	Template        string   `json:"template"`
	DatesExpression []string `json:"dates"`
}

func main() {
	log.Printf("log:START SendNotification")
	dynaClient = *db.InitalizeDynamoClient()
	lambda.Start(handler)
}
func handler(context context.Context, event MyEvent) {
	webhookURL := os.Getenv("SlackWebhookURl")

	var content string
	content, _ = handlers.CreateContent(dynaClient, event.DatesExpression, event.User, event.Template, event.ItemType)
	if err != nil {
		return
	}

	if content != "" {
		notification.SendNotificationToSlack(webhookURL, content)
	}
}
