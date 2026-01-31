package assert

import (
	"errors"
	"fmt"
)

func Condition(condition bool, failMessage string, failMessageParams ...any) {
	if false == condition {
		// The GO compiler does not allow using variadic arguments
		// from the function parameters, for Printf, Errorf, etc...
		// To work around this, we need to store the variadic arguments
		// in a local slice so the GO compiler thinks its safer
		var params []any = failMessageParams
		if len(params) > 0 {
			panic(fmt.Errorf(failMessage, params...))
		}
		panic(errors.New(failMessage))
	}
}
