package assert

import (
	"fmt"

	"github.com/JudicaelT/betterstandards/types"
)

func GreaterThan[T types.Numeric](a, b T) {
	if a <= b {
		panic(fmt.Errorf(
			"Failed asserting that value A ('%v') is greater than value B ('%v')",
			a,
			b,
		))
	}
}
