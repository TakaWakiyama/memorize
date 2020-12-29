package handlers

import (
	"github.com/guregu/dynamo"
	"memos/memos/pkg/memos"
)

const tableName = "Memos"

var (
	err error
)

func GetMemo(dynaClient *dynamo.DB, pk *string) interface {
	// dynamo.Table
	table := dynaClient.Table(tableName)
	return memos.Get(dynamo.Table, pk)
}
