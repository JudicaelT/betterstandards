package assert

import (
	"errors"
	"reflect"
)

func NotNil(value any) {
	if value == nil {
		panic(errors.New("Failed asserting that value is not nil"))
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
			panic(errors.New("Failed asserting that value is not nil"))
		}
	}
}
