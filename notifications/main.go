package main

// event から ユーザー名, テーブル名を取得 , 経過日数の配列, テンプレートの取得
// 経過日数(相対)を日付に直す
// 日付毎に batch get で取得したアイテム (最大50 25件)を　テンプレートをもとにテキストシリアライズ
// slack の webhook urlにpost リクエストを送信する
import (
	"context"
	"errors"
	"fmt"
	"log"
	"memos/common/db"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/guregu/dynamo"

	"memos/memos/pkg/memos"
	"memos/notifications/pkg/builder"
	"memos/notifications/pkg/notification"
)

var (
	dynaClient dynamo.DB
	err        error
)

func main() {
	log.Printf("log:START SendNotification")
	dynaClient = *db.InitalizeDynamoClient()
	lambda.Start(handler)
}

// MyEvent is passed from CluodWatch
type MyEvent struct {
	User            string   `json:"user"`
	ItemType        string   `json:"item_type"`
	Template        string   `json:"template"`
	DatesExpression []string `json:"dates"`
}

func handler(context context.Context, event MyEvent) {
	table := dynaClient.Table("Memos")
	webhookURL := os.Getenv("SlackWebhookURl")
	var result string
	dates := getDates(event.DatesExpression)
	for _, date := range dates {
		result += "---- " + date + " ----\n"
		memos := memos.GetMemos(table, event.User, event.ItemType, date)
		for h, memo := range memos {
			s, _ := builder.Parse(event.Template, memo.Detail)
			result += fmt.Sprintf("%d: %s\n", h+1, s)
		}
	}
	fmt.Println(result)
	result += "aaaaaaaa"
	if result == "" {
		notification.SendNotificationToSlack(webhookURL, result)
	}
}

var re, _ = regexp.Compile("([YMWDymwd])([1-9]?[0-9])")

func getDate(t time.Time) string {
	const layout2 = "2006-01-02"
	fmt.Println(t.Format(layout2))
	return t.Format(layout2)
}

// ArgDates is passed Time.AddDates
type ArgDates struct {
	Year  int
	Month int
	Day   int
}

func (d *ArgDates) setArg(dType string, digits int) error {
	lowwerDType := strings.ToLower(dType)
	if lowwerDType == "y" {
		d.Year += digits
	} else if lowwerDType == "m" {
		d.Month += digits
	} else if lowwerDType == "w" {
		d.Day += 7 * digits
	} else if lowwerDType == "d" {
		d.Day += digits
	} else {
		return errors.New("Invalid type")
	}
	return nil
}

func (d *ArgDates) setArgFromExp(exp string) error {
	matches := re.FindStringSubmatch(exp)
	errorString := "Unexpected error occured"
	if len(matches) != 3 {
		return errors.New(errorString)
	}
	dType, digitsString := matches[1], matches[2]
	digits, err := strconv.Atoi(digitsString)
	if err != nil {
		return errors.New(errorString)
	}
	err = d.setArg(dType, digits)
	if err != nil {
		return errors.New(errorString)
	}
	return nil
}

func getDates(dateExpressions []string) []string {
	var dates []string
	now := time.Now()
	for _, expression := range dateExpressions {
		a := ArgDates{}
		results := re.FindAllString(expression, 4)
		for _, exp := range results {
			err = a.setArgFromExp(exp)
			if err != nil {
				break
			}
		}
		if err != nil {
			continue
		}
		targetDate := now.AddDate(a.Year, a.Month, a.Day)
		const dateFormat = "2006-01-02"
		dates = append(dates, targetDate.Format(dateFormat))
	}
	return dates
}
