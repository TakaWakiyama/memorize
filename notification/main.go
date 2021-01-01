package main

// event から ユーザー名, テーブル名を取得 , 経過日数の配列, テンプレートの取得
// 経過日数(相対)を日付に直す
// 日付毎に batch get で取得したアイテム (最大50 25件)を　テンプレートをもとにテキストシリアライズ
// slack の webhook urlにpost リクエストを送信する
import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"

	"memos/notification/pkg"
)

const webhookURL = "https://hooks.slack.com/services/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"

// var (err error)

func main() {
	log.Printf("log:START Lambda Function")
	lambda.Start(handler)
}

// MyEvent is passed from CluodWatch
type MyEvent struct {
	User  string   `json:"user"`
	Table string   `json:"table"`
	Dates []string `json:"dates"`
}

func handler(context context.Context, event MyEvent) {
	eventJSON, _ := json.MarshalIndent(event, "", "  ")
	log.Printf("EVENT: %s", eventJSON)
	pkg.SendNotificationToSlack(webhookURL, "sample")
}
