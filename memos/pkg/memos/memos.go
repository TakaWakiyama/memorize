// Package memos is db layer
package memos

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/guregu/dynamo"
	"time"
)

var (
	err error
)

// Memo is DB schema
type Memo struct {
	User        string            `dynamo:"User,hash"`
	MemoId      string            `dynamo:"MemoId,range"`
	MemoType    string            `dynamo:"MemoType,range" json:"memo_type"`
	Detail      map[string]string `dynamo:"Detail" json:"detail"`
	DateCreated string            `dynamo:"DateCreated"`
}

// Get Single Item
func Get(table dynamo.Table, pk string) (*Memo, error) {
	fmt.Print("called Get")
	var memo Memo
	err := table.Get("memoID", pk).One(&memo)
	if err != nil {
		return nil, err
	}
	return &memo, nil
}

// Create is create memo
func Create(table dynamo.Table, memo *Memo) (string, error) {
	memo.MemoId = uuid.NewString()
	memo.DateCreated = getCurrentDate()
	if err := table.Put(memo).If("attribute_not_exists(MemoId)").Run(); err != nil {
		fmt.Printf("Failed to put item[%v]\n", err)
		return "", err
	}
	return memo.MemoId, nil
}

// GetMemos is filtering Memo with user and memotype
func GetMemos(table dynamo.Table, user, MemoType string) []Memo {
	var result []Memo
	err := table.Get("User", "Twaki").Filter("'MemoType' = ?", MemoType).All(&result)
	if err != nil {
		fmt.Printf("%v", err)
		return nil
	}
	return result
}

func getCurrentDate() string {
	t := time.Now()
	const layout2 = "2006-01-02"
	fmt.Println(t.Format(layout2))
	return t.Format(layout2)
}
