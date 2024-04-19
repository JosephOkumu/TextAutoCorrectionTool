package autocorrectiontool

import (
	"regexp"
	"strconv"
)

func HexToDecimal(str string) string {
	pattern := regexp.MustCompile(`(\b[0-9A-Fa-f]+)\s*\(\s*hex\s*\)`)
	// This line replaces occurrences of the pattern.
	// ReplaceAllStringFunc replaces matches of the pattern with the result of calling the provided function
	result := pattern.ReplaceAllStringFunc(str, func(match string) string {
		// This line extracts the hexadecimal word
		hexWrd := pattern.FindStringSubmatch(match)[1]
		// This line converts hexadecimal string to a decimal integer.
		decInt, err := strconv.ParseInt(hexWrd, 16, 64)
		// If hexword is not a valid hexadecimal string return the original matched string.
		if err != nil {
			return match
		}
		// This line converts a decimal integer back to a string representation
		return strconv.FormatInt(decInt, 10)
	})
	return result
}
