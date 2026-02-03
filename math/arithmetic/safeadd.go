package arithmetic

import (
	"fmt"
	"reflect"

	"github.com/JudicaelT/betterstandards/types"
)

func SafeAdd[T types.Number](a, b T, moreNumbersToAdd ...T) (sum T, overflowErr error) {
	sum = a + b
	if hasOverflowed(sum, a, b) {
		overflowErr = makeOverflowError(a, b)
	}
	for _, number := range moreNumbersToAdd {
		var sumTmp T = sum
		sum += number
		if overflowErr == nil && hasOverflowed(sum, sumTmp, number) {
			overflowErr = makeOverflowError(sumTmp, number)
		}
	}
	return
}

func hasOverflowed[T types.Number](sum, a, b T) bool {
	return (sum > a) != (b > 0)
}

func makeOverflowError[T types.Number](a, b T) error {
	return fmt.Errorf(
		"Value of type %s has overflowed when adding %v with %v",
		reflect.TypeOf(a),
		a,
		b,
	)
}
