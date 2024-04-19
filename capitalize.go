package autocorrectiontool

import (
	"strings"
	"regexp"
	
)

func Capitalize(str string) string {
	pattern := regexp.MustCompile(`(\b\w+)\s*\(\s*cap\s*\)`)
	// Replaces the pattern in the string with the result of the function
	result := pattern.ReplaceAllStringFunc(str, func(match string) string {
		// This line finds the match to group 1
		toCap := pattern.FindStringSubmatch(match)[1]
		// This line capitalizes the word before cap
		cap := strings.Title(toCap)
		return cap
	})
	return result
}