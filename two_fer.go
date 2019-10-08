// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

// ShareWith string -> string
func ShareWith(name string) string {
	// "One for X, one for me."
	// fmt.Println("name", name)

	// name == empty? default name of "you"
	if len(name) == 0 {
		name = "you"
	}

	str := "One for " + name + ", one for me."
	return str
}
