package arithmetic

import "github.com/JudicaelT/betterstandards/types"

func SafeAdd[T types.Number](a, b T, moreNumbersToAdd ...T) (sum T, hasOverflowed bool) {
	sum = a + b
	hasOverflowed = addHasOverflowed(sum, a, b)
	for _, number := range moreNumbersToAdd {
		var sumTmp T = sum
		sum += number
		hasOverflowed = hasOverflowed || addHasOverflowed(sum, sumTmp, number)
	}
	return
}

func addHasOverflowed[T types.Number](sum, a, b T) bool {
	return (sum > a) != (b > 0)
}
