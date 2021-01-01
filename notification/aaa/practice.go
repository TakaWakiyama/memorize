package main

import (
	//"encoding/json"
	"fmt"
)

var (
	err error
)

//MyEvent is sample struct
type MyEvent struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Dates []string `json:"dates"`
}

func main() {
	event := MyEvent{
		ID: 1,
		Dates: []string{
			"2020",
			"2021",
		},
	}
	handler(event)
}

func handler(event MyEvent) {
	fmt.Printf("ID is %d.Name is %s", event.ID, event.Name)
	fmt.Printf("%s", event.Dates)
}
