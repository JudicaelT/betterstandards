package slice

import (
	math "github.com/JudicaelT/betterstandards/math/arithmetic"
	"github.com/JudicaelT/betterstandards/types"
)

func Sum[T types.Number](slice []T) (T, error) {
	var sliceLen int = len(slice)
	if sliceLen < 2 {
		if sliceLen == 1 {
			return slice[0], nil
		}
		return 0, nil
	}
	return math.SafeAdd(slice[0], slice[1], slice[2:]...)
}
