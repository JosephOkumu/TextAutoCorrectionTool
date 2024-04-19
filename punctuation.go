package autocorrectiontool

import (
	"regexp"
	"strings"
)

func SetPunctuation(input string) string {
	punctPattern := regexp.MustCompile(`(\s*)([.,;:?!]+)(\s*)`)
	quotePattern := regexp.MustCompile(`'(\s*)(.*?)(\s*)'`)

	// Check for punctuation pattern match
	if punctPattern.MatchString(input) {
		input = punctPattern.ReplaceAllString(input, "$2 ")
	}

	// Check for quotation pattern match
	if quotePattern.MatchString(input){
		input = strings.TrimSpace(quotePattern.ReplaceAllString(input, "'$2'"))
	}
	return input
}