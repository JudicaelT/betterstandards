package slice_test

import (
	"math"
	"testing"

	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	"github.com/JudicaelT/betterstandards/math/arithmetic"
	"github.com/JudicaelT/betterstandards/slice"
	"github.com/stretchr/testify/assert"
)

func BenchmarkDiff(bench *testing.B) {
	intSlice := []int8{42, 21}
	codeUnderTest := func() { slice.Diff(intSlice) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkDiffWithOverflow(bench *testing.B) {
	intSlice := []int32{math.MinInt32, 1}
	codeUnderTest := func() { slice.Diff(intSlice) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func TestDiff(t *testing.T) {
	// Given a slice of integers
	intSlice := []int8{42, 21}

	// When we calculate the difference of all elements in the slice
	var diff int8
	var err error
	diff, err = slice.Diff(intSlice)

	// Then we should get the difference of all elements in the slice
	assert.Equal(t, int8(21), diff)
	// And it should not have overflowed
	assert.NoError(t, err)
}

func TestDiffWithEmptySlice(t *testing.T) {
	// Given an empty slice
	intSlice := []int{}

	// When we calculate the difference of the slice
	var diff int
	var err error
	diff, err = slice.Diff(intSlice)

	// Then we should get 0
	assert.Equal(t, 0, diff)
	// And it should not have overflowed
	assert.NoError(t, err)
}

func TestDiffWithSliceContainingOneElement(t *testing.T) {
	// Given a slice containing only one element
	intSlice := []int{42}

	// When we calculate the difference of the slice
	var diff int
	var err error
	diff, err = slice.Diff(intSlice)

	// Then we should get the only element in the slice
	assert.Equal(t, 42, diff)
	// And it should not have overflowed
	assert.NoError(t, err)
}

func TestDiffCausingOverflow(t *testing.T) {
	// Given a slice of int32 where subtracting the 2 elements
	// together causes an overflow
	intSlice := []int32{math.MinInt32, 1}

	// When we calculate the difference of all the elements in the slice
	var diff int32
	var err error
	diff, err = slice.Diff(intSlice)

	// Then we should get the difference of all the elements in the slice though it
	// should not correspond to the "real difference" because it should have overflowed
	assert.Equal(t, int32(2147483647), diff)
	// And we should get a hint that it overflowed
	assert.Same(t, arithmetic.SubOverflowErr, err)
}
