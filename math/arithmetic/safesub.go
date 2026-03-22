package arithmetic

import (
	"errors"

	"github.com/JudicaelT/betterstandards/types"
)

var SubOverflowErr error = errors.New("An overflow occurred while subtracting two numbers")

func SafeSub[T types.Number](a, b T, moreNumbersToSub ...T) (T, error) {
	var diff T = a - b
	var hasOverflowed bool = subHasOverflowed(diff, a, b)
	for _, number := range moreNumbersToSub {
		var diffTmp T = diff
		diff -= number
		hasOverflowed = hasOverflowed || subHasOverflowed(diff, diffTmp, number)
	}
	if hasOverflowed {
		return diff, SubOverflowErr
	}
	return diff, nil
}

func subHasOverflowed[T types.Number](diff, a, b T) bool {
	return (diff < a) != (b > 0)
}
