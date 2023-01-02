// Package raindrops takes an integer and returns its raindrop interpretation.
package raindrops

import (
	"fmt"
	"math"
)

// Convert returns a raindrop ('Pling', 'Plang', 'Plong') for provided integer input.
func Convert(a int) string {

	out := ""
	if math.Remainder(float64(a), 3) == 0 {
		out += "Pling"
	}
	if math.Remainder(float64(a), 5) == 0 {
		out += "Plang"
	}
	if math.Remainder(float64(a), 7) == 0 {
		out += "Plong"
	}
	if len(out) == 0 {
		out = fmt.Sprint(a)
	}
	return out
}
