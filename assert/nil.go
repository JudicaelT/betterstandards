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

	// nil pointers/chans/etc.... won't get caught by the first condition
	// because of interface conversion. Therefore, we have to use reflection
	// to check if the value is nil.
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case
		reflect.Chan, reflect.Func,
		reflect.Interface, reflect.Map,
		reflect.Pointer, reflect.Slice, reflect.UnsafePointer:
		if v.IsNil() {
			return
		}
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
