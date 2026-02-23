package slice

import (
	math "github.com/JudicaelT/betterstandards/math/arithmetic"
	"github.com/JudicaelT/betterstandards/types"
)

func Sum[T types.Number](slice []T) (sum T, hasOverflowed bool) {
	var sliceLen int = len(slice)
	if sliceLen < 2 {
		if sliceLen == 1 {
			return slice[0], false
		}
		return 0, false
	}
	return math.SafeAdd(slice[0], slice[1], slice[2:]...)
}
