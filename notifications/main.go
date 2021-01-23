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

	"memos/notifications/pkg/notification"
)

// const webhookURL = os.Getenv("SlackWebhookURl") || "https://hooks.slack.com/services/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"
const webhookURL = "https://hooks.slack.com/services/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"

var dynaClient dynamo.DB

type Memo struct {
	User     string `dynamo:"User,hash"`
	MemoId   string `dynamo:"MemoId,range"`
	MemoType string `dynamo:"MemoType,range"`
}

func main() {
	log.Printf("log:START SendNotification")
	dynaClient = *db.InitalizeDynamoClient()
	lambda.Start(handler)
}

// MyEvent is passed from CluodWatch
type MyEvent struct {
	User     string   `json:"user"`
	ItemType string   `json:"item_type"`
	Dates    []string `json:"dates"`
}

func handler(context context.Context, event MyEvent) {
	result := getMemos(event.User, event.ItemType)
	fmt.Printf(result)
	if result != "" {
		notification.SendNotificationToSlack(webhookURL, result)
	}
}

func getMemos(user, MemoType string) string {
	var result []Memo
	err := dynaClient.Table("Memos").Get("User", "Twaki").Filter("'MemoType' = ?", MemoType).All(&result)
	if err != nil {
		fmt.Printf("%v", err)
		return ""
	}
	//item := dynaClient.Table("Items").Get("User", "Twaki")
	fmt.Print(result)
	// fmt.Println(item, result, user, itemType)
	return ""
}
