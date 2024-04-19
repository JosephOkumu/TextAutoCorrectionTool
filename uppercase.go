package autocorrectiontool

import (
	"strings"
	"regexp"
)

func UpperCase(str string) string {
	pattern := regexp.MustCompile(`(\b\w+)\s*\(\s*up\s*\)`)
	result := pattern.ReplaceAllStringFunc(str, func(match string) string {
		toUp := pattern.FindStringSubmatch(match)[1]
		uP := strings.ToUpper(toUp)
		return uP
	})
	return result
}