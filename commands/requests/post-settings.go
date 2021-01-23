package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const url = "https://foa17sekhl.execute-api.ap-northeast-1.amazonaws.com/Prod/settings"

type Data struct {
	User              string `json:"user"`
	Category          string `json:"category"`
	IsActive          int    `json:"is_active"`
	ExecutionInterval []int  `json:"execution_interval"`
	TimeExecute       string `json:"time_execute"`
	Template          string `json:"template"`
}

func main() {
	data := Data{
		User:              "Twaki",
		Category:          "Url",
		IsActive:          1,
		ExecutionInterval: []int{1, 2, 3},
		TimeExecute:       "17:00",
		Template:          `<a href="{url}">url</a>`,
	}
	body, _ := json.Marshal(data)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v", res.Body)
}
