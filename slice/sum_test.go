package slice_test

import (
	"math"
	"testing"

	"github.com/JudicaelT/betterstandards/slice"
)

func TestSum(t *testing.T) {
	// Given a slice of integers
	intSlice := []int8{1, 2, 3, 4}

	// When we calculate the sum of all elements in the slice
	var sum int8
	var err error
	sum, err = slice.Sum(intSlice)

	// Then we should get the sum of all elements in the slice
	if sum != 10 {
		t.Error("Failed asserting that the sum of the slice is 10. Got", sum)
	}

	// And there should not be an error
	if err != nil {
		t.Error("Failed asserting that that there is no error. Got ", err.Error())
	}
}

func TestSumWithEmptySlice(t *testing.T) {
	// Given an empty slice
	intSlice := []int{}

	// When we calculate the sum of the slice
	var sum int
	var err error
	sum, err = slice.Sum(intSlice)

	// Then we should get 0
	if sum != 0 {
		t.Error("Failed asserting that the sum of the slice is 0. Got", sum)
	}

	// And there should not be an error
	if err != nil {
		t.Error("Failed asserting that that there is no error. Got ", err.Error())
	}
}

func TestSumWithSliceContainingOneElement(t *testing.T) {
	// Given a slice containing only one element
	intSlice := []int{42}

	// When we calculate the sum of the slice
	var sum int
	var err error
	sum, err = slice.Sum(intSlice)

	// Then we should get the only element in the slice
	if sum != 42 {
		t.Error("Failed asserting that the sum of the slice is 42. Got", sum)
	}

	// And there should not be an error
	if err != nil {
		t.Error("Failed asserting that that there is no error. Got ", err.Error())
	}
}

func TestSumCausingOverflow(t *testing.T) {
	// Given a slice of int32 where adding the 2 elements
	// together causes an overflow
	intSlice := []int32{math.MaxInt32, 1}

	// When we calculate the sum of all the elements in the slice
	var sum int32
	var err error
	sum, err = slice.Sum(intSlice)

	// Then we should get the sum of all the elements in the slice though it
	// should not correspond to the "real sum" because it should have overflowed
	if sum != -2147483648 {
		t.Error("Failed asserting that the sum of the slice is -2147483648. Got:", sum)
	}

	// But we should also get an error warning us that an overflow occured
	if err == nil {
		t.Error("Failed asserting that error is not nil")
	}
	expectedMessage := "Value of type int32 has overflowed when adding 2147483647 with 1"
	actualMessage := err.Error()
	if expectedMessage != actualMessage {
		t.Errorf(
			"Expected error message '%s' but got '%s'",
			expectedMessage,
			actualMessage,
		)
	}
}
