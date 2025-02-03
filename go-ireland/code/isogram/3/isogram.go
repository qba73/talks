package isogram

import (
	"strings"
)

// START OMIT
func IsIsogram(word string) bool {

	word = strings.ToUpper(word)

	for i, c := range word {
		if c == ' ' || c == '-' { // HL
			continue
		}
		for j := i + 1; j < len(word); j++ {
			if word[i] == word[j] {
				return false
			}
		}
	}
	return true
}

// END OMIT
