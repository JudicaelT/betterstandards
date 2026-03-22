package arithmetic_test

import (
	"math"
	"testing"

	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	"github.com/JudicaelT/betterstandards/math/arithmetic"
	"github.com/stretchr/testify/assert"
)

func BenchmarkSafeSub(bench *testing.B) {
	var a, b int8 = 10, 2
	codeUnderTest := func() { arithmetic.SafeSub(a, b) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkSafeSubWithOverflow(bench *testing.B) {
	var a, b int8 = math.MinInt8, 1
	codeUnderTest := func() { arithmetic.SafeSub(a, b) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkSafeSubWithVariaticArguments(bench *testing.B) {
	var a, b, c int8 = 44, 1, 1
	codeUnderTest := func() { arithmetic.SafeSub(a, b, c) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkSafeSubWithVariaticArgumentsCausingOverflow(bench *testing.B) {
	var a, b, c int8 = 1, 3, math.MaxInt8
	codeUnderTest := func() { arithmetic.SafeSub(a, b, c) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func TestSafeSub(t *testing.T) {
	// Given 2 integers
	var a, b int8 = 10, 2

	// When we substract them together using SafeSub
	var diff int8
	var err error
	diff, err = arithmetic.SafeSub(a, b)

	// Then we should get the difference of those 2 integers
	assert.Equal(t, int8(8), diff)
	// And it should not have overflowed
	assert.NoError(t, err)
}

func TestSafeSubWithOverflow(t *testing.T) {
	// Given 2 integers that will overflow if subtracted together
	var a, b int8 = math.MinInt8, 1

	// When we sub them together using SafeSub
	var diff int8
	var err error
	diff, err = arithmetic.SafeSub(a, b)

	// Then we should get the diff of those 2 integers though it
	// should not correspond to the "real diff" because it should have overflowed
	assert.Equal(t, int8(127), diff)
	// And we should get a hint that it overflowed
	assert.Same(t, arithmetic.SubOverflowErr, err)
}

func TestSafeSubWithVariaticArguments(t *testing.T) {
	// Given 3 integers
	var a, b, c int8 = 44, 1, 1

	// When we sub them together using SafeSub
	var diff int8
	var err error
	diff, err = arithmetic.SafeSub(a, b, c)

	// Then we should get the diff of those 3 integers
	assert.Equal(t, int8(42), diff)
	// And it should not have overflowed
	assert.NoError(t, err)
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
	assert.Equal(t, int8(127), diff)
	// And we should get a hint that it overflowed
	assert.Same(t, arithmetic.SubOverflowErr, err)
}
