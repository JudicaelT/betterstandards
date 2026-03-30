package arithmetic_test

import (
	"math"
	"testing"

	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	"github.com/JudicaelT/betterstandards/math/arithmetic"
	"github.com/stretchr/testify/assert"
)

func BenchmarkSafeDiv(bench *testing.B) {
	var a, b int8 = 42, 2
	codeUnderTest := func() { arithmetic.SafeDiv(a, b) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkSafeDivWithOverflow(bench *testing.B) {
	var a, b int8 = math.MinInt8, -1
	codeUnderTest := func() { arithmetic.SafeDiv(a, b) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkSafeDivByZero(bench *testing.B) {
	var a, b int8 = 42, 0
	codeUnderTest := func() { arithmetic.SafeDiv(a, b) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkSafeDivWithVariaticArguments(bench *testing.B) {
	var a, b, c int8 = 42, 2, 2
	codeUnderTest := func() { arithmetic.SafeDiv(a, b, c) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkSafeDivWithVariaticArgumentsCausingOverflow(bench *testing.B) {
	var a, b, c int8 = math.MinInt8, -1, 2
	codeUnderTest := func() { arithmetic.SafeDiv(a, b, c) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func BenchmarkSafeDivByZeroWithVariaticArguments(bench *testing.B) {
	var a, b, c int8 = 42, 2, 0
	codeUnderTest := func() { arithmetic.SafeDiv(a, b, c) }
	benchmark.AvgRuntime(bench, codeUnderTest)
	benchmark.AssertNoAllocs(bench, codeUnderTest)
}

func TestSafeDiv(t *testing.T) {
	{
		// Given 2 int
		var a, b int = 42, 2
		// When we divide them together using SafeDiv
		var quotient int
		var err error
		quotient, err = arithmetic.SafeDiv(a, b)
		// Then we should get the quotient of those 2 integers
		assert.Equal(t, 21, quotient)
		// And it should not have overflowed
		assert.NoError(t, err)
	}
	{
		// Same with int8
		var a, b int8 = 42, 2
		quotient, err := arithmetic.SafeDiv(a, b)
		assert.Equal(t, int8(21), quotient)
		assert.NoError(t, err)
	}
	{
		// Same with int16
		var a, b int16 = 42, 2
		quotient, err := arithmetic.SafeDiv(a, b)
		assert.Equal(t, int16(21), quotient)
		assert.NoError(t, err)
	}
	{
		// Same with int32
		var a, b int32 = 42, 2
		quotient, err := arithmetic.SafeDiv(a, b)
		assert.Equal(t, int32(21), quotient)
		assert.NoError(t, err)
	}
	{
		// Same with int64
		var a, b int64 = 42, 2
		quotient, err := arithmetic.SafeDiv(a, b)
		assert.Equal(t, int64(21), quotient)
		assert.NoError(t, err)
	}
	{
		// Same with uint
		var a, b uint = 42, 2
		quotient, err := arithmetic.SafeDiv(a, b)
		assert.Equal(t, uint(21), quotient)
		assert.NoError(t, err)
	}
	{
		// Same with float32
		var a, b float32 = 42, 2
		quotient, err := arithmetic.SafeDiv(a, b)
		assert.Equal(t, float32(21), quotient)
		assert.NoError(t, err)
	}
}

func TestSafeDivWithOverflow(t *testing.T) {
	{
		// Given 2 int
		var a, b int = math.MinInt, -1
		// When we divide them together using SafeDiv
		var quotient int
		var err error
		quotient, err = arithmetic.SafeDiv(a, b)
		// Then we should get the quotient of those 2 integers though it
		// should not correspond to the "real quotient" because it should have overflowed
		assert.Equal(t, math.MinInt, quotient)
		// And we should get a hint that it overflowed
		assert.Same(t, arithmetic.DivOverflowErr, err)
	}
	{
		// Same with int8
		var a, b int8 = math.MinInt8, -1
		quotient, err := arithmetic.SafeDiv(a, b)
		assert.Equal(t, int8(math.MinInt8), quotient)
		assert.Same(t, arithmetic.DivOverflowErr, err)
	}
	{
		// Same with int16
		var a, b int16 = math.MinInt16, -1
		quotient, err := arithmetic.SafeDiv(a, b)
		assert.Equal(t, int16(math.MinInt16), quotient)
		assert.Same(t, arithmetic.DivOverflowErr, err)
	}
	{
		// Same with int32
		var a, b int32 = math.MinInt32, -1
		quotient, err := arithmetic.SafeDiv(a, b)
		assert.Equal(t, int32(math.MinInt32), quotient)
		assert.Same(t, arithmetic.DivOverflowErr, err)
	}
	{
		// Same with int64
		var a, b int64 = math.MinInt64, -1
		quotient, err := arithmetic.SafeDiv(a, b)
		assert.Equal(t, int64(math.MinInt64), quotient)
		assert.Same(t, arithmetic.DivOverflowErr, err)
	}
}

func TestSafeDivByZero(t *testing.T) {
	// Given 2 int
	var a, b int = 42, 0
	// When we divide them together using SafeDiv
	var quotient int
	var err error
	quotient, err = arithmetic.SafeDiv(a, b)
	// Then we should get zero because B equals zero
	assert.Equal(t, 0, quotient)
	// And there should be an error
	assert.Same(t, arithmetic.DivByZeroErr, err)
}

func TestSafeDivWithVariaticArguments(t *testing.T) {
	{
		// Given 3 int
		var a, b, c int = 42, 2, 2
		// When we divide them together using SafeDiv
		var quotient int
		var err error
		quotient, err = arithmetic.SafeDiv(a, b, c)
		// Then we should get the quotient of those 3 integers
		// (note that it gets floored automatically because it is an int)
		assert.Equal(t, 10, quotient)
		// And it should not have overflowed
		assert.NoError(t, err)
	}
	{
		// Same with int8
		var a, b, c int8 = 42, 2, 2
		quotient, err := arithmetic.SafeDiv(a, b, c)
		assert.Equal(t, int8(10), quotient)
		assert.NoError(t, err)
	}
	{
		// Same with int16
		var a, b, c int16 = 42, 2, 2
		quotient, err := arithmetic.SafeDiv(a, b, c)
		assert.Equal(t, int16(10), quotient)
		assert.NoError(t, err)
	}
	{
		// Same with int32
		var a, b, c int32 = 42, 2, 2
		quotient, err := arithmetic.SafeDiv(a, b, c)
		assert.Equal(t, int32(10), quotient)
		assert.NoError(t, err)
	}
	{
		// Same with int64
		var a, b, c int64 = 42, 2, 2
		quotient, err := arithmetic.SafeDiv(a, b, c)
		assert.Equal(t, int64(10), quotient)
		assert.NoError(t, err)
	}
	{
		// Same with uint
		var a, b, c uint = 42, 2, 2
		quotient, err := arithmetic.SafeDiv(a, b, c)
		assert.Equal(t, uint(10), quotient)
		assert.NoError(t, err)
	}
	{
		// Same with float
		var a, b, c float32 = 42, 2, 2
		quotient, err := arithmetic.SafeDiv(a, b, c)
		assert.Equal(t, float32(10.5), quotient)
		assert.NoError(t, err)
	}
}

func TestSafeDivWithVariaticArgumentsCausingOverflow(t *testing.T) {
	{
		// Given 3 int
		var a, b, c int = math.MinInt, -1, 2
		// When we divide them together using SafeDiv
		var quotient int
		var err error
		quotient, err = arithmetic.SafeDiv(a, b, c)
		// Then we should get the quotient of those 3 integers though it
		// should not correspond to the "real quotient" because it should have overflowed
		assert.Equal(t, math.MinInt/2, quotient)
		// And we should get a hint that it overflowed
		assert.Same(t, arithmetic.DivOverflowErr, err)
	}
	{
		// Same with int8
		var a, b, c int8 = math.MinInt8, -1, 2
		quotient, err := arithmetic.SafeDiv(a, b, c)
		assert.Equal(t, int8(math.MinInt8/2), quotient)
		assert.Same(t, arithmetic.DivOverflowErr, err)
	}
	{
		// Same with int16
		var a, b, c int16 = math.MinInt16, -1, 2
		quotient, err := arithmetic.SafeDiv(a, b, c)
		assert.Equal(t, int16(math.MinInt16/2), quotient)
		assert.Same(t, arithmetic.DivOverflowErr, err)
	}
	{
		// Same with int32
		var a, b, c int32 = math.MinInt32, -1, 2
		quotient, err := arithmetic.SafeDiv(a, b, c)
		assert.Equal(t, int32(math.MinInt32/2), quotient)
		assert.Same(t, arithmetic.DivOverflowErr, err)
	}
	{
		// Same with int64
		var a, b, c int64 = math.MinInt64, -1, 2
		quotient, err := arithmetic.SafeDiv(a, b, c)
		assert.Equal(t, int64(math.MinInt64/2), quotient)
		assert.Same(t, arithmetic.DivOverflowErr, err)
	}
}

func TestSafeDivByZeroWithVariaticArguments(t *testing.T) {
	// Given 3 int
	var a, b, c int = 42, 2, 0
	// When we divide them together using SafeDiv
	var quotient int
	var err error
	quotient, err = arithmetic.SafeDiv(a, b, c)
	// Then we should get zero because C equals zero
	assert.Equal(t, 0, quotient)
	// And there should be an error
	assert.Same(t, arithmetic.DivByZeroErr, err)
}
