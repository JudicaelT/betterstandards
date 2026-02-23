package slice_test

import (
	"math"
	"testing"

	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	"github.com/JudicaelT/betterstandards/slice"
)

func BenchmarkSum(bench *testing.B) {
	intSlice := []int8{1, 2, 3, 4}
	codeUnderTest := func() { slice.Sum(intSlice) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkSumWithOverflow(bench *testing.B) {
	intSlice := []int32{math.MaxInt32, 1}
	codeUnderTest := func() { slice.Sum(intSlice) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func TestSum(t *testing.T) {
	// Given a slice of integers
	intSlice := []int8{1, 2, 3, 4}

	// When we calculate the sum of all elements in the slice
	var sum int8
	var hasOverflowed bool
	sum, hasOverflowed = slice.Sum(intSlice)

	// Then we should get the sum of all elements in the slice
	if sum != 10 {
		t.Error("Failed asserting that the sum of the slice is 10. Got", sum)
	}

	// And it should not have overflowed
	if hasOverflowed {
		t.Error("Failed asserting that there was no overflow")
	}
}

func TestSumWithEmptySlice(t *testing.T) {
	// Given an empty slice
	intSlice := []int{}

	// When we calculate the sum of the slice
	var sum int
	var hasOverflowed bool
	sum, hasOverflowed = slice.Sum(intSlice)

	// Then we should get 0
	if sum != 0 {
		t.Error("Failed asserting that the sum of the slice is 0. Got", sum)
	}

	// And it should not have overflowed
	if hasOverflowed {
		t.Error("Failed asserting that there was no overflow")
	}
}

func TestSumWithSliceContainingOneElement(t *testing.T) {
	// Given a slice containing only one element
	intSlice := []int{42}

	// When we calculate the sum of the slice
	var sum int
	var hasOverflowed bool
	sum, hasOverflowed = slice.Sum(intSlice)

	// Then we should get the only element in the slice
	if sum != 42 {
		t.Error("Failed asserting that the sum of the slice is 42. Got", sum)
	}

	// And it should not have overflowed
	if hasOverflowed {
		t.Error("Failed asserting that there was no overflow")
	}
}

func TestSumCausingOverflow(t *testing.T) {
	// Given a slice of int32 where adding the 2 elements
	// together causes an overflow
	intSlice := []int32{math.MaxInt32, 1}

	// When we calculate the sum of all the elements in the slice
	var sum int32
	var hasOverflowed bool
	sum, hasOverflowed = slice.Sum(intSlice)

	// Then we should get the sum of all the elements in the slice though it
	// should not correspond to the "real sum" because it should have overflowed
	if sum != -2147483648 {
		t.Error("Failed asserting that the sum of the slice is -2147483648. Got:", sum)
	}

	// And we should get a hint that it overflowed
	if !hasOverflowed {
		t.Error("Failed asserting that slice.Sum() caused an overflow")
	}
}
