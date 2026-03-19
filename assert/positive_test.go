package assert_test

import (
	"testing"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
)

func BenchmarkAssertPositive(b *testing.B) {
	codeUnderTest := func() { assert.Positive(42) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func TestAssertPositive(t *testing.T) {
	// assert.Positive() should not panic
	functionUnderTest := "assert.Positive"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When positive value is passed
	assert.Positive(0)
}

func TestAssertPositiveWithNegativeNumber(t *testing.T) {
	// assert.Positive() should panic
	functionUnderTest := "assert.Positive"
	expectedErrorMessage := "Failed asserting that value is positive. Got: -0.00021"
	defer test.ShouldPanic(t, functionUnderTest, expectedErrorMessage)

	// When value negative value is passed
	assert.Positive(float64(-0.00021))
}
