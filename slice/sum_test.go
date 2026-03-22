package slice_test

import (
	"math"
	"testing"

	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	"github.com/JudicaelT/betterstandards/math/arithmetic"
	"github.com/JudicaelT/betterstandards/slice"
	"github.com/stretchr/testify/assert"
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
	var err error
	sum, err = slice.Sum(intSlice)

	// Then we should get the sum of all elements in the slice
	assert.Equal(t, int8(10), sum)
	// And it should not have overflowed
	assert.NoError(t, err)
}

func TestSumWithEmptySlice(t *testing.T) {
	// Given an empty slice
	intSlice := []int{}

	// When we calculate the sum of the slice
	var sum int
	var err error
	sum, err = slice.Sum(intSlice)

	// Then we should get 0
	assert.Equal(t, 0, sum)
	// And it should not have overflowed
	assert.NoError(t, err)
}

func TestSumWithSliceContainingOneElement(t *testing.T) {
	// Given a slice containing only one element
	intSlice := []int{42}

	// When we calculate the sum of the slice
	var sum int
	var err error
	sum, err = slice.Sum(intSlice)

	// Then we should get the only element in the slice
	assert.Equal(t, 42, sum)
	// And it should not have overflowed
	assert.NoError(t, err)
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
	assert.Equal(t, int32(-2147483648), sum)
	// And we should get a hint that it overflowed
	assert.Same(t, arithmetic.AddOverflowErr, err)
}
