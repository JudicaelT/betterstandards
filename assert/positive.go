package assert

import (
	"fmt"

	"github.com/JudicaelT/betterstandards/types"
)

func Positive[T types.Numeric](value T) {
	if value < 0 {
		panic(fmt.Errorf(
			"Failed asserting that value is positive. Got: %v",
			value,
		))
	}
}
