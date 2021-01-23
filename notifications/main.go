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
	"regexp"

	"memos/notifications/pkg/notification"
)

var dynaClient dynamo.DB

type Memo struct {
	User     string            `dynamo:"User,hash"`
	MemoId   string            `dynamo:"MemoId,range"`
	MemoType string            `dynamo:"MemoType,range"`
	Detail   map[string]string `dynamo:"Detail"`
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
	Template string   `json:"template"`
	Dates    []string `json:"dates"`
}

func handler(context context.Context, event MyEvent) {
	memos := getMemos(event.User, event.ItemType)
	webhookURL := os.Getenv("SlackWebhookURl")
	var result string
	for i, memo := range memos {
		s, _ := Parse(event.Template, memo.Detail)
		result += fmt.Sprintf("%d: %s\n", i+1, s)
	}
	if result != "" {
		notification.SendNotificationToSlack(webhookURL, result)
	}
}

func getMemos(user, MemoType string) []Memo {
	var result []Memo
	err := dynaClient.Table("Memos").Get("User", "Twaki").Filter("'MemoType' = ?", MemoType).All(&result)
	if err != nil {
		fmt.Printf("%v", err)
		return nil
	}
	return result
}

// Parse is
func Parse(template string, various map[string]string) (string, error) {
	re := regexp.MustCompile(`\{[\s]{0,}([a-zA-Z]+)[\s]{0,}\}`)
	cb := func(s string) string {
		extractNames := re.FindStringSubmatch(s)
		if len(extractNames) != 2 {
			return ""
		}
		attributeName := extractNames[1]
		// fmt.Printf("tname %v", tname) output -> tname [{ word} word]tname [{ type } type]
		if result := various[attributeName]; result != "" {
			return result
		}
		return ""
	}
	result := re.ReplaceAllStringFunc(template, cb)
	return result, nil
}
