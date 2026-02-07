package assert

import "fmt"

func EmptySlice[T any](value []T) {
	var numberOfElements int = len(value)
	if numberOfElements != 0 {
		panic(fmt.Errorf(
			"Failed asserting that slice is empty. Found '%d' elements",
			numberOfElements,
		))
	}
}

func EmptyString(s string) {
	Condition(
		s == "",
		"Failed asserting that string is empty. Got '%s'",
		s,
	)
}

func EmptyMap[K comparable, T any](value map[K]T) {
	var numberOfElements int = len(value)
	Condition(
		numberOfElements == 0,
		"Failed asserting that map is empty. Found '%d' elements",
		numberOfElements,
	)
}

func EmptyChannel[T any](value chan T) {
	var numberOfElements int = len(value)
	Condition(
		numberOfElements == 0,
		"Failed asserting that channel is empty. Found '%d' elements",
		numberOfElements,
	)
}
