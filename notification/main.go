package main

// event から ユーザー名, テーブル名を取得 , 経過日数の配列, テンプレートの取得
// 経過日数(相対)を日付に直す
// 日付毎に batch get で取得したアイテム (最大50 25件)を　テンプレートをもとにテキストシリアライズ
// slack の webhook urlにpost リクエストを送信する
import (
	"context"
	"encoding/json"
	"log"

	// "github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// var (err error)

func main() {
	log.Printf("log:START Lambda Function")
	lambda.Start(handler)
}

type MyEvent struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handler(context context.Context, event MyEvent) {
	eventJson, _ := json.MarshalIndent(event, "", "  ")
	log.Printf("EVENT: %s", eventJson)
}
