package arithmetic

import "github.com/JudicaelT/betterstandards/types"

func SafeAdd[T types.Number](a, b T, moreNumbersToAdd ...T) (sum T, hasOverflowed bool) {
	sum = a + b
	hasOverflowed = (sum > a) != (b > 0)
	for _, number := range moreNumbersToAdd {
		var sumTmp T = sum
		sum += number
		if (false == hasOverflowed) {
			hasOverflowed = (sum > sumTmp) != (number > 0)
		}
	}
	return
}
