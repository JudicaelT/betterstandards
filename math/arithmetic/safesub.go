package arithmetic

import "github.com/JudicaelT/betterstandards/types"

func SafeSub[T types.Number](a, b T, moreNumbersToSub ...T) (diff T, hasOverflowed bool) {
	diff = a - b
	hasOverflowed = subHasOverflowed(diff, a, b)
	for _, number := range moreNumbersToSub {
		var diffTmp T = diff
		diff -= number
		hasOverflowed = hasOverflowed || subHasOverflowed(diff, diffTmp, number)
	}
	return
}

func subHasOverflowed[T types.Number](diff, a, b T) bool {
	return (diff < a) != (b > 0)
}
