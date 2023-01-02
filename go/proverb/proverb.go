// Package proverb returns a rhyme based in the input keywords
package proverb

import (
	"fmt"
)

// Proverb takes a list of keywords as input and returns a list of relevant rhymes
func Proverb(rhyme []string) []string {

	var out []string
	for i := 1; i < len(rhyme); i++ {
		r := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i])
		out = append(out, r)
	}

	if len(rhyme) > 0 {
		r := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		out = append(out, r)
	}

	return out
}
