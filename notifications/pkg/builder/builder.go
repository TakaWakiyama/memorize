package builder

import (
	"regexp"
)

// Parse is
func Parse(template string, various map[string]string) (string, error) {
	re := regexp.MustCompile(`\{[\s]{0,}([a-zA-Z]+)[\s]{0,}\}`)
	cb := func(s string) string {
		extractNames := re.FindStringSubmatch(s)
		if len(extractNames) != 2 {
			return ""
		}
		attributeName := extractNames[1]
		// fmt.Printf("tname %v", tname) output -> tname [{ word} word]tname [{ type } type]
		if result := various[attributeName]; result != "" {
			return result
		}
		return ""
	}
	result := re.ReplaceAllStringFunc(template, cb)
	return result, nil
}
