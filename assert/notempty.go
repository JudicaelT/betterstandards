package assert

import "errors"

func NotEmptySlice[T any](slice []T) {
	if len(slice) == 0 {
		panic(errors.New("Failed asserting that the given slice is not empty"))
	}
}

func NotEmptyString(s string) {
	if len(s) == 0 {
		panic(errors.New("Failed asserting that the given string is not empty"))
	}
}

func NotEmptyMap[K comparable, T any](m map[K]T) {
	if len(m) == 0 {
		panic(errors.New("Failed asserting that the given map is not empty"))
	}
}

func NotEmptyChannel[T any](channel chan T) {
	if len(channel) == 0 {
		panic(errors.New("Failed asserting that the given channel is not empty"))
	}
}
