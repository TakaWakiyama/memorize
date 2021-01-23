package settings

import (
	"github.com/guregu/dynamo"
)

// Setting is db schema
type Setting struct {
	User              string `dynamo:"user"`     // HASH
	Category          string `dynamo:"category"` // Range
	IsActive          int    `dynamo:"is_active"`
	ExecutionInterval []int  `dynamo:"execution_interval"`
	TimeExecute       string `dynamo:"time_execute"`
	Template          string `dynamo:"template"`
}

//TableKey is Setting Key
type TableKey struct {
	User     string `json:"user"`
	Category string `json:"category"`
}

//Create is
func Create(table dynamo.Table, setting Setting) (*TableKey, error) {

	// auto increment にする必要がある
	if err := table.Put(setting).Run(); err != nil {
		return nil, err
	}
	tableKey := TableKey{
		User:     setting.User,
		Category: setting.Category,
	}
	return &tableKey, nil
}
