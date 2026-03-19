package assert

import (
	"fmt"

	"github.com/JudicaelT/betterstandards/types"
)

func Negative[T types.Number](value T) {
	if value >= 0 {
		panic(fmt.Errorf(
			"Failed asserting that value is negative. Got: %v",
			value,
		))
	}
}
