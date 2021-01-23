package builder

import (
	"fmt"
	"regexp"
)

var (
	err error
)

// Build is
func Build(template string, memos map[string]interface{}) string {
	text, err := parse(template, memos)
	if err != nil {
		return fmt.Sprintf("%s \n%v", template, err)
	}
	return text
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
