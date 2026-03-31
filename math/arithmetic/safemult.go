package arithmetic

import (
	"errors"

	"github.com/JudicaelT/betterstandards/types"
)

var MultOverflowErr error = errors.New("An overflow occurred while multiplying two numbers")

func SafeMult[T types.Numeric](a, b T, moreNumbersToMult ...T) (T, error) {
	var product T = a * b
	var hasOverflowed bool = multHasOverflowed(product, a, b)
	for _, number := range moreNumbersToMult {
		var productTmp T = product
		product *= number
		hasOverflowed = hasOverflowed || multHasOverflowed(product, productTmp, number)
	}
	if hasOverflowed {
		return product, MultOverflowErr
	}
	return product, nil
}

func multHasOverflowed[T types.Numeric](product, a, b T) bool {
	return a != 0 && product/a != b
}
