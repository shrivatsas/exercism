// Package twofer returns a two for one.
// One for `name` and one for me.
package twofer

import "fmt"

// ShareWith takes a name and returns a two-fer string.
// The input defaults to 'you', if input is empty.
func ShareWith(name string) string {
	var you = "you"
	if name != "" {
		you = name
	}
	return fmt.Sprintf("One for %s, one for me.", you)
}
