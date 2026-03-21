package assert

import "fmt"

func SliceLen[T any](expectedLen int, slice []T) {
	var actualLen int = len(slice)
	if actualLen != expectedLen {
		panic(fmt.Errorf(
			"Failed asserting that length of slice is %d. Got %d",
			expectedLen,
			actualLen,
		))
	}
}

func StringLen(expectedLen int, str string) {
	var actualLen int = len(str)
	if actualLen != expectedLen {
		panic(fmt.Errorf(
			"Failed asserting that length of string is %d. Got %d",
			expectedLen,
			actualLen,
		))
	}
}

func MapLen[K comparable, T any](expectedLen int, m map[K]T) {
	var actualLen int = len(m)
	if actualLen != expectedLen {
		panic(fmt.Errorf(
			"Failed asserting that length of map is %d. Got %d",
			expectedLen,
			actualLen,
		))
	}
}

func ChannelLen[T any](expectedLen int, channel chan T) {
	var actualLen int = len(channel)
	if actualLen != expectedLen {
		panic(fmt.Errorf(
			"Failed asserting that length of channel is %d. Got %d",
			expectedLen,
			actualLen,
		))
	}
}
