package test

import (
	"reflect"
	"testing"
)

func ShouldPanic(t *testing.T, functionUnderTest, expectedMessage string) {
	r := recover()
	if r == nil {
		t.Errorf(
			"%s() should have panicked but did not",
			functionUnderTest,
		)
		return
	}

	err, isError := r.(error)
	if !isError {
		t.Errorf(
			"%s() should have panicked with value of type 'error' but panicked with '%s' instead",
			functionUnderTest,
			reflect.TypeOf(r),
		)
		return
	}

	actualMessage := err.Error()
	if actualMessage != expectedMessage {
		t.Errorf(
			"Error message '%s' does not match with expected '%s'",
			actualMessage,
			expectedMessage,
		)
	}
}
