package assert

import "errors"

func NotNil(value any) {
	if value == nil {
		panic(errors.New("Failed asserting that value is not nil"))
	}
}
