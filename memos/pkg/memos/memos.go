// Package memo is db layer
package memo

import (
	"fmt"

	"github.com/guregu/dynamo"
)

var (
	err error
)

// Memo is DB schema
type Memo struct {
	memoID   string
	memoType string
	value    string
}

// Get Single Item
func Get(table *dynamo.Table, pk string) (*Memo, error) {
	fmt.Print("called Get")
	var memo Memo
	err := table.Get("memoID", pk).One(&memo)
	if err != nil {
		return nil, err
	}
	return &memo, nil
}
