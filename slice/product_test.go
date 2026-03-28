package slice_test

import (
	"math"
	"testing"

	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	"github.com/JudicaelT/betterstandards/math/arithmetic"
	"github.com/JudicaelT/betterstandards/slice"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProduct(bench *testing.B) {
	intSlice := []int8{21, 2}
	codeUnderTest := func() { slice.Product(intSlice) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkProductWithOverflow(bench *testing.B) {
	intSlice := []int32{math.MinInt32, -1}
	codeUnderTest := func() { slice.Product(intSlice) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func TestProduct(t *testing.T) {
	// Given a slice of integers
	intSlice := []int8{21, 2}

	// When we calculate the product of all elements in the slice
	var diff int8
	var err error
	diff, err = slice.Product(intSlice)

	// Then we should get the product of all elements in the slice
	assert.Equal(t, int8(42), diff)
	// And it should not have overflowed
	assert.NoError(t, err)
}

func TestProductWithEmptySlice(t *testing.T) {
	// Given an empty slice
	intSlice := []int{}

	// When we calculate the product of the slice
	var diff int
	var err error
	diff, err = slice.Product(intSlice)

	// Then we should get 0
	assert.Equal(t, 0, diff)
	// And it should not have overflowed
	assert.NoError(t, err)
}

func TestProductWithSliceContainingOneElement(t *testing.T) {
	// Given a slice containing only one element
	intSlice := []int{42}

	// When we calculate the product of the slice
	var diff int
	var err error
	diff, err = slice.Product(intSlice)

	// Then we should get the only element in the slice
	assert.Equal(t, 42, diff)
	// And it should not have overflowed
	assert.NoError(t, err)
}

func TestProductCausingOverflow(t *testing.T) {
	// Given a slice of int32 where multiplying the 2 elements
	// together causes an overflow
	intSlice := []int32{math.MinInt32, -1}

	// When we calculate the product of all the elements in the slice
	var diff int32
	var err error
	diff, err = slice.Product(intSlice)

	// Then we should get the product of all the elements in the slice though it
	// should not correspond to the "real product" because it should have overflowed
	assert.Equal(t, int32(-2147483648), diff)
	// And we should get a hint that it overflowed
	assert.Same(t, arithmetic.MultOverflowErr, err)
}
