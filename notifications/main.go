package main

// event から ユーザー名, テーブル名を取得 , 経過日数の配列, テンプレートの取得
// 経過日数(相対)を日付に直す
// 日付毎に batch get で取得したアイテム (最大50 25件)を　テンプレートをもとにテキストシリアライズ
// slack の webhook urlにpost リクエストを送信する
import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/guregu/dynamo"
	"log"
	"memos/common/db"
	"os"

	"memos/memos/pkg/memos"
	"memos/notifications/pkg/builder"
	"memos/notifications/pkg/notification"
)

var dynaClient dynamo.DB

func main() {
	log.Printf("log:START SendNotification")
	dynaClient = *db.InitalizeDynamoClient()
	lambda.Start(handler)
}

// MyEvent is passed from CluodWatch
type MyEvent struct {
	User     string   `json:"user"`
	ItemType string   `json:"item_type"`
	Template string   `json:"template"`
	Dates    []string `json:"dates"`
}

func handler(context context.Context, event MyEvent) {
	table := dynaClient.Table("Memos")
	memos := memos.GetMemos(table, event.User, event.ItemType)
	webhookURL := os.Getenv("SlackWebhookURl")
	var result string
	for i, memo := range memos {
		s, _ := builder.Parse(event.Template, memo.Detail)
		result += fmt.Sprintf("%d: %s\n", i+1, s)
	}
	if result != "" {
		notification.SendNotificationToSlack(webhookURL, result)
	}
}
