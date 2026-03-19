package assert

import (
	"fmt"

	"github.com/JudicaelT/betterstandards/types"
)

func LessThan[T types.Number](a, b T) {
	if a >= b {
		panic(fmt.Errorf(
			"Failed asserting that value A ('%v') is less than value B ('%v')",
			a,
			b,
		))
	}
}
