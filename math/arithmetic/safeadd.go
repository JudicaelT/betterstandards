package arithmetic

import (
	"errors"

	"github.com/JudicaelT/betterstandards/types"
)

var AddOverflowErr error = errors.New("An overflow occurred while adding two numbers")

func SafeAdd[T types.Numeric](a, b T, moreNumbersToAdd ...T) (T, error) {
	var sum T = a + b
	var hasOverflowed bool = addHasOverflowed(sum, a, b)
	for _, number := range moreNumbersToAdd {
		var sumTmp T = sum
		sum += number
		hasOverflowed = hasOverflowed || addHasOverflowed(sum, sumTmp, number)
	}
	if hasOverflowed {
		return sum, AddOverflowErr
	}
	return sum, nil
}

func addHasOverflowed[T types.Numeric](sum, a, b T) bool {
	return (sum > a) != (b > 0)
}
