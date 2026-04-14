package assert

import "errors"

func Ok[T any](value T, ok bool) T {
	if !ok {
		panic(errors.New("Failed asserting that ok is true"))
	}
	return value
}
