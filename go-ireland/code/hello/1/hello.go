package hello

// START OMIT
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me." // HL
}

// END OMIT
