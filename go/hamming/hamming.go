// Package hamming takes 2 DNA strands as strings and calculates the point difference between them.
package hamming

import (
	"errors"
)

// Distance takes 2 strings as input and returns the distance for equal length strands.
func Distance(a, b string) (int, error) {

	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		return 0, errors.New("strands are not of same length")
	}

	distance := 0
	for i := range ar {
		if ar[i] != br[i] {
			distance++
		}
	}
	return distance, nil
}
