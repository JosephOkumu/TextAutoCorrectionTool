package autocorrectiontool

import (
	"regexp"
	"strings"
)

func LowerCase(str string) string {
	pattern := regexp.MustCompile(`(\b\w+)\s*\(\s*low\s*\)`)
	result := pattern.ReplaceAllStringFunc(str, func(match string) string {
		toLow := pattern.FindStringSubmatch(match)[1]
		low := strings.ToLower(toLow)
		return low
	})
	return result
}
