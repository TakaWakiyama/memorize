// Package memos is db layer
package memos

import (
	"fmt"

	"github.com/guregu/dynamo"
)

var (
	err error
)

// Memo is DB schema
type Memo struct {
	MemoID   string
	MemoType string
	Value    string
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
	fmt.Print("called Put")
	if err := table.Put(memo).Run(); err != nil {
		fmt.Printf("Failed to put item[%v]\n", err)
		return "", err
	}

	return memo.MemoID, nil
}
