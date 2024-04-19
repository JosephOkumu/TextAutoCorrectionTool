package autocorrectiontool

import (
	"strings"
	"strconv"
)

func ModifyText(input string) string {
	words := strings.Fields(input)
	for i := 0; i < len(words); i++ {
		if strings.HasSuffix(words[i], "(up,") {
			words = modifyWords(words, i, func(s string) string {
				return strings.ToUpper(s)
			})
		}else if strings.HasSuffix(words[i], "(low,") {
			words = modifyWords(words, i, func(s string) string {
				return strings.ToLower(s)
			})
		}else if strings.HasSuffix(words[i], "(cap,") {
			words = modifyWords(words, i, func(s string) string {
				return strings.Title(s)
			})
		}
	}
	return strings.Join(words, " ")
}

func modifyWords(words []string, i int, modifier func(string)string) []string {
	if i+1 >= len(words) {
		return words
	}
	count, err := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
	if err != nil {
		return words
	}
	if count > i {
		count = i
	}
	for j := 1; j <= count; j++ {
		words[i-j] = modifier(words[i-j])
	}
	return append(words[:i], words[i+2:]...)
}