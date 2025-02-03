package isogram

import "unicode"

// START OMIT
func IsIsogram(word string) bool {

	letters := make(map[rune]bool) // HL

	for _, l := range word {
		letter := unicode.ToLower(l)
		if letter == '-' || unicode.IsSpace(letter) { // HL
			continue
		}
		if letters[letter] {
			return false
		}
		letters[letter] = true
	}
	return true
}

// END OMIT
