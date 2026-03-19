package assert

import (
	"fmt"

	"github.com/JudicaelT/betterstandards/types"
)

func LessThanEq[T types.Number](a, b T) {
	if a > b {
		panic(fmt.Errorf(
			"Failed asserting that value A ('%v') is equal or less than value B ('%v')",
			a,
			b,
		))
	}
}
