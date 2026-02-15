package assert

import "errors"

func True(value bool) {
	if value == false {
		panic(errors.New("Failed asserting that value is true"))
	}
}
