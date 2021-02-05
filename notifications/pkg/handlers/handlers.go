package handlers

import (
	"errors"
	"fmt"
	"memos/memos/pkg/memos"
	"memos/notifications/pkg/builder"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/guregu/dynamo"
)

var (
	err error
)

// CreateContent is create notification content from event
func CreateContent(dynaClient dynamo.DB, dateExpressions []string, userName string, template string, itemType string) (string, error) {
	var content string
	table := dynaClient.Table("Memos")
	dates := getDates(dateExpressions)
	for _, date := range dates {
		content += "---- " + date + " ----\n"
		memos := memos.GetMemos(table, userName, itemType, date)
		for h, memo := range memos {
			s, _ := builder.Parse(template, memo.Detail)
			content += fmt.Sprintf("%d: %s\n", h+1, s)
		}
	}
	return content, nil
}

var re, _ = regexp.Compile("([YMWDymwd])([1-9]?[0-9])")

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
		targetDate := now.AddDate(-a.Year, -a.Month, -a.Day)
		const dateFormat = "2006-01-02"
		dates = append(dates, targetDate.Format(dateFormat))
	}
	return dates
}
