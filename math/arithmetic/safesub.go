package arithmetic

import (
	"fmt"
	"reflect"

	"github.com/JudicaelT/betterstandards/types"
)

func SafeSub[T types.Number](a, b T, moreNumbersToSub ...T) (diff T, overflowErr error) {
	diff = a - b
	if subHasOverflowed(diff, a, b) {
		overflowErr = makeSubOverflowError(a, b)
	}
	for _, number := range moreNumbersToSub {
		var diffTmp T = diff
		diff -= number
		if overflowErr == nil && subHasOverflowed(diff, diffTmp, number) {
			overflowErr = makeSubOverflowError(diffTmp, number)
		}
	}
	return
}

func subHasOverflowed[T types.Number](diff, a, b T) bool {
	return (diff < a) != (b > 0)
}

func makeSubOverflowError[T types.Number](a, b T) error {
	return fmt.Errorf(
		"Value of type %s has overflowed when subtracting %v with %v",
		reflect.TypeOf(a),
		a,
		b,
	)
}
