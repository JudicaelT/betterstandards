package arithmetic_test

import (
	"math"
	"testing"

	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	"github.com/JudicaelT/betterstandards/math/arithmetic"
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
	var hasOverflowed bool
	sum, hasOverflowed = arithmetic.SafeAdd(a, b)

	// Then we should get the sum of those 2 integers
	if sum != 4 {
		t.Error("Failed asserting that the sum is equal to 4. Got:", sum)
	}

	// And it should not have overflowed
	if hasOverflowed {
		t.Error("Failed asserting addition did not cause an overflow")
	}
}

func TestSafeAddWithOverflow(t *testing.T) {
	// Given 2 integers that will overflow if you add them together
	var a, b int8 = math.MaxInt8, 1

	// When we add them together using SafeAdd
	var sum int8
	var hasOverflowed bool
	sum, hasOverflowed = arithmetic.SafeAdd(a, b)

	// Then we should get the sum of those 2 integers though it
	// should not correspond to the "real sum" because it should have overflowed
	if sum != -128 {
		t.Error("Failed asserting that the sum is equal to -128. Got", sum)
	}

	// And we should get a hint that it overflowed
	if !hasOverflowed {
		t.Error("Failed asserting addition overflowed")
	}
}

func TestSafeAddWithVariaticArguments(t *testing.T) {
	// Given 3 integers
	var a, b, c int8 = 1, 1, 40

	// When we add them together using SafeAdd
	var sum int8
	var hasOverflowed bool
	sum, hasOverflowed = arithmetic.SafeAdd(a, b, c)

	// Then we should get the sum of those 3 integers
	if sum != 42 {
		t.Error("Failed asserting that the sum is equal to 42. Got:", sum)
	}

	// And it should not have overflowed
	if hasOverflowed {
		t.Error("Failed asserting addition did not cause an overflow")
	}
}

func TestSafeAddWithVariaticArgumentsCausingOverflow(t *testing.T) {
	// Given 3 integers
	var a, b, c int8 = 1, 1, math.MaxInt8

	// When we add them together using SafeAdd
	var sum int8
	var hasOverflowed bool
	sum, hasOverflowed = arithmetic.SafeAdd(a, b, c)

	// Then we should get the sum of those 3 integers though it
	// should not correspond to the "real sum" because it should have overflowed
	if sum != -127 {
		t.Error("Failed asserting that the sum is equal to -127. Got", sum)
	}

	// And we should get a hint that it overflowed
	if !hasOverflowed {
		t.Error("Failed asserting addition overflowed")
	}
}
