package slice_test

import (
	"math"
	"testing"

	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	"github.com/JudicaelT/betterstandards/math/arithmetic"
	"github.com/JudicaelT/betterstandards/slice"
	"github.com/stretchr/testify/assert"
)

func BenchmarkQuotient(bench *testing.B) {
	intSlice := []int8{42, 2}
	codeUnderTest := func() { slice.Quotient(intSlice) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkQuotientWithOverflow(bench *testing.B) {
	intSlice := []int32{math.MinInt32, -1}
	codeUnderTest := func() { slice.Quotient(intSlice) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkQuotientWithZero(bench *testing.B) {
	intSlice := []int32{42, 0}
	codeUnderTest := func() { slice.Quotient(intSlice) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func TestQuotient(t *testing.T) {
	// Given a slice of integers
	intSlice := []int8{42, 2}

	// When we calculate the quotient of all elements in the slice
	var diff int8
	var err error
	diff, err = slice.Quotient(intSlice)

	// Then we should get the quotient of all elements in the slice
	assert.Equal(t, int8(21), diff)
	// And it should not have overflowed
	assert.NoError(t, err)
}

func TestQuotientWithEmptySlice(t *testing.T) {
	// Given an empty slice
	intSlice := []int{}

	// When we calculate the quotient of the slice
	var diff int
	var err error
	diff, err = slice.Quotient(intSlice)

	// Then we should get 0
	assert.Equal(t, 0, diff)
	// And it should not have overflowed
	assert.NoError(t, err)
}

func TestQuotientWithSliceContainingOneElement(t *testing.T) {
	// Given a slice containing only one element
	intSlice := []int{42}

	// When we calculate the quotient of the slice
	var diff int
	var err error
	diff, err = slice.Quotient(intSlice)

	// Then we should get the only element in the slice
	assert.Equal(t, 42, diff)
	// And it should not have overflowed
	assert.NoError(t, err)
}

func TestQuotientCausingOverflow(t *testing.T) {
	// Given a slice of int32 where dividing the 2 elements
	// together causes an overflow
	intSlice := []int32{math.MinInt32, -1}

	// When we calculate the quotient of all the elements in the slice
	var diff int32
	var err error
	diff, err = slice.Quotient(intSlice)

	// Then we should get the quotient of all the elements in the slice though it
	// should not correspond to the "real quotient" because it should have overflowed
	assert.Equal(t, int32(-2147483648), diff)
	// And we should get a hint that it overflowed
	assert.Same(t, arithmetic.DivOverflowErr, err)
}

func TestQuotientWithZero(t *testing.T) {
	// Given a slice of integers
	intSlice := []int{42, 0}

	// When we calculate the quotient of all elements in the slice
	var diff int
	var err error
	diff, err = slice.Quotient(intSlice)

	// Then we should get the quotient of all elements in the slice
	assert.Equal(t, 0, diff)
	// And there should be an error
	assert.Same(t, arithmetic.DivByZeroErr, err)
}
