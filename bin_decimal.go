package autocorrectiontool

import (
	"regexp"
	"strconv"
)

func BinToDecimal(str string) string {
	pattern := regexp.MustCompile(`(\b[01]+)\s*\(\s*bin\s*\)`)
	// ReplaceAllStringFunc replaces matches of the pattern with the result of calling the provided function. 
	result := pattern.ReplaceAllStringFunc(str, func(match string) string {
		// This line extracts the binary word.
		binWrd := pattern.FindStringSubmatch(match)[1]
		// This line converts binary string to decimal integer
		decInt, err := strconv.ParseInt(binWrd, 2, 64)
		// if binWrd is not a valid binary string return the original matched string
		if err != nil {
			return match
		}
		// This line converts a decimal integer back to a string representation
		return strconv.FormatInt(decInt, 10)
	})
	return result
}
