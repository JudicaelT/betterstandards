package arithmetic_test

import (
	"math"
	"testing"

	"github.com/JudicaelT/betterstandards/math/arithmetic"
)

func TestSafeSub(t *testing.T) {
	// Given 2 integers
	var a, b int8 = 10, 2

	// When we substract them together using SafeSub
	var diff int8
	var err error
	diff, err = arithmetic.SafeSub(a, b)

	// Then we should get the difference of those 2 integers
	if diff != 8 {
		t.Error("Failed asserting that the difference is equal to 8. Got:", diff)
	}

	// And there should not be an error
	if err != nil {
		t.Error("Failed asserting that there is no error. Got:", err)
	}
}

func TestSafeSubWithOverflow(t *testing.T) {
	// Given 2 integers that will overflow if you sub them together
	var a, b int8 = math.MinInt8, 1

	// When we sub them together using SafeSub
	var diff int8
	var err error
	diff, err = arithmetic.SafeSub(a, b)

	// Then we should get the diff of those 2 integers though it
	// should not correspond to the "real diff" because it should have overflowed
	if diff != 127 {
		t.Error("Failed asserting that the difference is equal to 127. Got", diff)
	}

	// And there should be an error
	if err == nil {
		t.Error("Failed asserting that there was an error")
		return
	}
	expectedMessage := "Value of type int8 has overflowed when subtracting -128 with 1"
	actualMessage := err.Error()
	if expectedMessage != actualMessage {
		t.Errorf(
			"Expected error message '%s' but got '%s'",
			expectedMessage,
			actualMessage,
		)
	}
}

func TestSafeSubWithVariaticArguments(t *testing.T) {
	// Given 3 integers
	var a, b, c int8 = 44, 1, 1

	// When we sub them together using SafeSub
	var diff int8
	var err error
	diff, err = arithmetic.SafeSub(a, b, c)

	// Then we should get the diff of those 3 integers
	if diff != 42 {
		t.Error("Failed asserting that the difference is equal to 42. Got:", diff)
	}

	// And there should not be an error
	if err != nil {
		t.Error("Failed asserting that there is no error. Got:", err)
	}
}

func TestSafeSubWithVariaticArgumentsCausingOverflow(t *testing.T) {
	// Given 3 integers
	var a, b, c int8 = 1, 3, math.MaxInt8

	// When we add them together using SafeSub
	var diff int8
	var err error
	diff, err = arithmetic.SafeSub(a, b, c)

	// Then we should get the diff of those 3 integers though it
	// should not correspond to the "real diff" because it should have overflowed
	if diff != 127 {
		t.Error("Failed asserting that the difference is equal to 127. Got", diff)
	}

	// And there should be an error
	if err == nil {
		t.Error("Failed asserting that err not nil")
		return
	}
	expectedMessage := "Value of type int8 has overflowed when subtracting -2 with 127"
	actualMessage := err.Error()
	if expectedMessage != actualMessage {
		t.Errorf(
			"Expected error message '%s' but got '%s'",
			expectedMessage,
			actualMessage,
		)
	}
}
