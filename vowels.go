package autocorrectiontool

import (
	"strings"
)

func ReplaceWithAn(str string) string {
	// split the string into a slice of strings.
	words := strings.Fields(str)
	// define a slice of byte containing vowels both uppercase and small
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'h', 'A', 'E', 'I', 'O', 'U', 'H'}
	laString := ""

	for i, word := range words {
		for _, v := range vowels {
			// check if word is "a" and if the next word's first letter is a vowel
			if word == "a" && words[i+1][0] == v {
				words[i] = "an"
			// check if word is "A" and if the next word's first letter is a vowel
			}else if word == "A" && words[i+1][0] == v {
				words[i] = "An"
			}
		}
		// append the words from  the changed "a" and after, to the words before changed "a" 
		words = append(words[:i], words[i:]...)
	}
	// transforming the slice of words into a string
	laString = strings.Join(words, " ")
	return laString
}