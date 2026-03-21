package arithmetic

import "github.com/JudicaelT/betterstandards/types"

func SafeMult[T types.Number](a, b T, moreNumbersToMult ...T) (product T, hasOverflowed bool) {
	product = a * b
	hasOverflowed = multHasOverflowed(product, a, b)
	for _, number := range moreNumbersToMult {
		var productTmp T = product
		product *= number
		hasOverflowed = hasOverflowed || multHasOverflowed(product, productTmp, number)
	}
	return
}

func multHasOverflowed[T types.Number](product, a, b T) bool {
	return a != 0 && product/a != b
}
