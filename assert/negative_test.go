package assert_test

import (
	"testing"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
)

func BenchmarkAssertNegative(b *testing.B) {
	codeUnderTest := func() { assert.Negative(-42) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func TestAssertNegative(t *testing.T) {
	// assert.Negative() should not panic
	functionUnderTest := "assert.Negative"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When a negative value is passed
	assert.Negative(float64(-0.00021))
}

func TestAssertNegativeWithPositiveNumber(t *testing.T) {
	// assert.Negative() should panic
	functionUnderTest := "assert.Negative"
	expectedErrorMessage := "Failed asserting that value is negative. Got: 0"
	defer test.ShouldPanic(t, functionUnderTest, expectedErrorMessage)

	// When a positive value is passed
	assert.Negative(0)
}
