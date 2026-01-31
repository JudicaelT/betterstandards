package test

import "testing"

func ShouldNotPanic(t *testing.T, functionUnderTest string) {
	r := recover()
	if r != nil {
		t.Errorf("%s() should not have panicked: %v", functionUnderTest, r)
	}
}
