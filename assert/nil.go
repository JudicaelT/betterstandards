package assert

import (
	"errors"
	"fmt"
	"reflect"
)

func Nil(value any) {
	if value == nil {
		return
	}

	var errorMessage string
	if err, isError := value.(error); isError {
		errorMessage = fmt.Sprintf(
			"Expected value to be nil. Got error: '%s'",
			err.Error(),
		)
	} else {
		errorMessage = fmt.Sprintf(
			"Expected value to be nil. Got '%s'",
			reflect.TypeOf(value),
		)
	}
	panic(errors.New(errorMessage))
}
