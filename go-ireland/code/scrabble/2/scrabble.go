package scrabble

import "strings"

// START OMIT
var points = [26]int{
	1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3,
	1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10,
}

func Score(word string) int {
	sum := 0
	workString := strings.ToUpper(word)

	for _, set := range workString {
		sum += points[set-65]
	}
	return sum
}

// END OMIT
