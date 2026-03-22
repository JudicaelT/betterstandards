package arithmetic_test

import (
	"math"
	"testing"

	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	"github.com/JudicaelT/betterstandards/math/arithmetic"
	"github.com/stretchr/testify/assert"
)

func BenchmarkSafeAdd(bench *testing.B) {
	var a, b int8 = 2, 2
	codeUnderTest := func() { arithmetic.SafeAdd(a, b) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkSafeAddWithOverflow(bench *testing.B) {
	var a, b int8 = math.MaxInt8, 1
	codeUnderTest := func() { arithmetic.SafeAdd(a, b) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkSafeAddWithVariaticArguments(bench *testing.B) {
	var a, b, c int8 = 1, 1, 40
	codeUnderTest := func() { arithmetic.SafeAdd(a, b, c) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkSafeAddWithVariaticArgumentsCausingOverflow(bench *testing.B) {
	var a, b, c int8 = 1, 1, math.MaxInt8
	codeUnderTest := func() { arithmetic.SafeAdd(a, b, c) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func TestSafeAdd(t *testing.T) {
	// Given 2 integers
	var a, b int8 = 2, 2

	// When we add them together using SafeAdd
	var sum int8
	var err error
	sum, err = arithmetic.SafeAdd(a, b)

	// Then we should get the sum of those 2 integers
	assert.Equal(t, int8(4), sum)
	// And it should not have overflowed
	assert.NoError(t, err)
}

func TestSafeAddWithOverflow(t *testing.T) {
	// Given 2 integers that will overflow if added together
	var a, b int8 = math.MaxInt8, 1

	// When we add them together using SafeAdd
	var sum int8
	var err error
	sum, err = arithmetic.SafeAdd(a, b)

	// Then we should get the sum of those 2 integers though it
	// should not correspond to the "real sum" because it should have overflowed
	assert.Equal(t, int8(-128), sum)
	// And we should get a hint that it overflowed
	assert.Same(t, arithmetic.AddOverflowErr, err)
}

func TestSafeAddWithVariaticArguments(t *testing.T) {
	// Given 3 integers
	var a, b, c int8 = 1, 1, 40

	// When we add them together using SafeAdd
	var sum int8
	var err error
	sum, err = arithmetic.SafeAdd(a, b, c)

	// Then we should get the sum of those 3 integers
	assert.Equal(t, int8(42), sum)
	// And it should not have overflowed
	assert.NoError(t, err)
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
	assert.Equal(t, int8(-127), sum)
	// And we should get a hint that it overflowed
	assert.Same(t, arithmetic.AddOverflowErr, err)
}
