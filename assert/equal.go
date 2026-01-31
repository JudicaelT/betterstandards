package assert

import "fmt"

func Equal[T comparable](a, b T) {
	if a != b {
		panic(fmt.Errorf("Failed asserting that '%v' is equal to '%v'", a, b))
	}
}
