package hello

import "fmt"

// START OMIT
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name) // HL
}

//END OMIT
