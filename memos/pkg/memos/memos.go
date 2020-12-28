// Package memo is db layer
package memo

import (
	"fmt"
)

// Memo is DB schema
type Memo struct {
	memoID   int
	memoType string
	value    string
}

func getMemos() int {
	return 1
}

func getMemo(pk int) int {
	return pk
}
