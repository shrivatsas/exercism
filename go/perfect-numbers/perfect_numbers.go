package perfect

import (
	"errors"
)

// Classification defines the class of the integer
type Classification string

// The type of Classifications
const (
	ClassificationAbundant  Classification = "ClassificationAbundant"
	ClassificationDeficient Classification = "ClassificationDeficient"
	ClassificationPerfect   Classification = "ClassificationPerfect"
)

// ErrOnlyPositive when non-positive numbers are added
var ErrOnlyPositive = errors.New("Not a positve number")

func findFactors(in int64) []int64 {
	var res []int64
	var i int64
	for i = 1; float64(i) <= float64(in)/2; i++ {
		if in%i == 0 {
			res = append(res, int64(i))
		}
	}
	return res
}

// Classify returns the type of number
func Classify(in int64) (Classification, error) {
	if in <= 0 {
		return Classification("None"), ErrOnlyPositive
	}

	s := int64(0)
	factors := findFactors(in)
	for _, val := range factors {
		s += val
	}

	if s < in {
		return ClassificationDeficient, nil
	} else if s == in {
		return ClassificationPerfect, nil
	} else {
		return ClassificationAbundant, nil
	}
}
