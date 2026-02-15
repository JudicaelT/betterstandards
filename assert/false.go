package assert

import "errors"

func False(value bool) {
	if value == true {
		panic(errors.New("Failed asserting that value is false"))
	}
}
