package assert

import "fmt"

func NotEqual[T comparable](a, b T) {
	if a == b {
		panic(fmt.Errorf("Failed asserting that value A and B ('%v') are not equal", a))
	}
}
