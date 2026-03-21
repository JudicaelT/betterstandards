package arithmetic_test

import (
	"testing"

	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	"github.com/JudicaelT/betterstandards/math/arithmetic"
	"github.com/stretchr/testify/assert"
)

func BenchmarkSafeMult(bench *testing.B) {
	var a, b int8 = 21, 2
	codeUnderTest := func() { arithmetic.SafeMult(a, b) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkSafeMultWithOverflow(bench *testing.B) {
	var a, b int8 = 64, 2
	codeUnderTest := func() { arithmetic.SafeMult(a, b) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkSafeMultWithVariaticArguments(bench *testing.B) {
	var a, b, c int8 = 21, 2, 2
	codeUnderTest := func() { arithmetic.SafeMult(a, b, c) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkSafeMultWithVariaticArgumentsCausingOverflow(bench *testing.B) {
	var a, b, c int8 = 1, 64, 2
	codeUnderTest := func() { arithmetic.SafeMult(a, b, c) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func TestSafeMult(t *testing.T) {
	// Given 2 integers
	var a, b int8 = 21, 2

	// When we multiply them together using SafeMult
	var product int8
	var hasOverflowed bool
	product, hasOverflowed = arithmetic.SafeMult(a, b)

	// Then we should get the product of those 2 integers
	assert.Equal(t, int8(42), product)
	// And it should not have overflowed
	assert.False(t, hasOverflowed)
}

func TestSafeMultWithOverflow(t *testing.T) {
	// Given 2 integers that will overflow if multiplied together
	var a, b int8 = 64, 2

	// When we multiply them together using SafeMult
	var product int8
	var hasOverflowed bool
	product, hasOverflowed = arithmetic.SafeMult(a, b)

	// Then we should get the product of those 2 integers though it
	// should not correspond to the "real product" because it should have overflowed
	assert.Equal(t, int8(-128), product)

	// And we should get a hint that it overflowed
	assert.True(t, hasOverflowed)
}

func TestSafeMultWithVariaticArguments(t *testing.T) {
	// Given 3 integers
	var a, b, c int8 = 21, 2, 2

	// When we multiply them together using SafeMult
	var product int8
	var hasOverflowed bool
	product, hasOverflowed = arithmetic.SafeMult(a, b, c)

	// Then we should get the product of those 3 integers
	assert.Equal(t, int8(84), product)

	// And it should not have overflowed
	assert.False(t, hasOverflowed)
}

func TestSafeMultWithVariaticArgumentsCausingOverflow(t *testing.T) {
	// Given 3 integers
	var a, b, c int8 = 1, 64, 2

	// When we multiply them together using SafeMult
	var product int8
	var hasOverflowed bool
	product, hasOverflowed = arithmetic.SafeMult(a, b, c)

	// Then we should get the product of those 3 integers though it
	// should not correspond to the "real product" because it should have overflowed
	assert.Equal(t, int8(-128), product)

	// And we should get a hint that it overflowed
	assert.True(t, hasOverflowed)
}
