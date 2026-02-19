package arithmetic

import (
	"fmt"
	"reflect"

	"github.com/JudicaelT/betterstandards/types"
)

func SafeAdd[T types.Number](a, b T, moreNumbersToAdd ...T) (sum T, overflowErr error) {
	sum = a + b
	if addHasOverflowed(sum, a, b) {
		overflowErr = makeAddOverflowError(a, b)
	}
	for _, number := range moreNumbersToAdd {
		var sumTmp T = sum
		sum += number
		if overflowErr == nil && addHasOverflowed(sum, sumTmp, number) {
			overflowErr = makeAddOverflowError(sumTmp, number)
		}
	}
	return
}

func addHasOverflowed[T types.Number](sum, a, b T) bool {
	return (sum > a) != (b > 0)
}

func makeAddOverflowError[T types.Number](a, b T) error {
	return fmt.Errorf(
		"Value of type %s has overflowed when adding %v with %v",
		reflect.TypeOf(a),
		a,
		b,
	)
}
