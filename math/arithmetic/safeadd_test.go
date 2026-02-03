package arithmetic_test

import (
	"math"
	"testing"

	"github.com/JudicaelT/betterstandards/math/arithmetic"
)

func TestSafeAdd(t *testing.T) {
	// Given 2 integers
	var a, b int8 = 2, 2

	// When we add them together using SafeAdd
	var sum int8
	var err error
	sum, err = arithmetic.SafeAdd(a, b)

	// Then we should get the sum of those 2 integers
	if sum != 4 {
		t.Error("Failed asserting that the sum is equal to 4. Got:", sum)
	}

	// And there should not be an error
	if err != nil {
		t.Error("Failed asserting that there is no error. Got:", err)
	}
}

func TestSafeAddWithOverflow(t *testing.T) {
	// Given 2 integers that will overflow if you add them together
	var a, b int8 = math.MaxInt8, 1

	// When we add them together using SafeAdd
	var sum int8
	var err error
	sum, err = arithmetic.SafeAdd(a, b)

	// Then we should get the sum of those 2 integers though it
	// should not correspond to the "real sum" because it should have overflowed
	if sum != -128 {
		t.Error("Failed asserting that the sum is equal to -128. Got", sum)
	}

	// And there should be an error
	if err == nil {
		t.Error("Failed asserting that there was an error")
		return
	}
	expectedMessage := "Value of type int8 has overflowed when adding 127 with 1"
	actualMessage := err.Error()
	if expectedMessage != actualMessage {
		t.Errorf(
			"Expected error message '%s' but got '%s'",
			expectedMessage,
			actualMessage,
		)
	}
}

func TestSafeAddWithVariaticArguments(t *testing.T) {
	// Given 3 integers
	var a, b, c int8 = 1, 1, 40

	// When we add them together using SafeAdd
	var sum int8
	var err error
	sum, err = arithmetic.SafeAdd(a, b, c)

	// Then we should get the sum of those 3 integers
	if sum != 42 {
		t.Error("Failed asserting that the sum is equal to 42. Got:", sum)
	}

	// And there should not be an error
	if err != nil {
		t.Error("Failed asserting that there is no error. Got:", err)
	}
}

func TestSafeAddWithVariaticArgumentsCausingOverflow(t *testing.T) {
	// Given 3 integers
	var a, b, c int8 = 1, 1, math.MaxInt8

	// When we add them together using SafeAdd
	var sum int8
	var err error
	sum, err = arithmetic.SafeAdd(a, b, c)

	// Then we should get the sum of those 3 integers though it
	// should not correspond to the "real sum" because it should have overflowed
	if sum != -127 {
		t.Error("Failed asserting that the sum is equal to -127. Got", sum)
	}

	// And there should be an error
	if err == nil {
		t.Error("Failed asserting that err not nil")
		return
	}
	expectedMessage := "Value of type int8 has overflowed when adding 2 with 127"
	actualMessage := err.Error()
	if expectedMessage != actualMessage {
		t.Errorf(
			"Expected error message '%s' but got '%s'",
			expectedMessage,
			actualMessage,
		)
	}
}
