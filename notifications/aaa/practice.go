package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

var (
	err error
)

func main() {
	// map[string]interface{}

	jsonBlob := []byte(`
		{
			"word": "invite",
			"type": "noun",
			"meaning": "make a polite, formal, or friendly request to (someone) to go somewhere or to do something.",
			"id": 1
		}
	`)
	var event interface{}
	err := json.Unmarshal(jsonBlob, &event)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	devent, ok := event.(map[string]interface{})
	if !ok {
		return
	}
	// to stringfy id
	template := "{word}{id}\n Word: { word}\n Type: { type }\n Meaning : {meaning }\n"
	text, _ := parse(template, devent)
	fmt.Println(text)
}

func parse(template string, various map[string]interface{}) (string, error) {
	re := regexp.MustCompile(`\{[\s]{0,}([a-z]+)[\s]{0,}\}`)
	cb := func(s string) string {
		extractNames := re.FindStringSubmatch(s)
		if len(extractNames) != 2 {
			return ""
		}
		attributeName := extractNames[1]
		// fmt.Printf("tname %v", tname) output -> tname [{ word} word]tname [{ type } type]
		if result := various[attributeName]; result != "" {
			return fmt.Sprintf("%v", result)
		}
		return ""
	}
	result := re.ReplaceAllStringFunc(template, cb)
	return result, nil
}
