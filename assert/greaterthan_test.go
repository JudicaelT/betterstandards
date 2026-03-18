package assert_test

import (
	"testing"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
)

func BenchmarkAssertGreaterThan(b *testing.B) {
	codeUnderTest := func() { assert.GreaterThan(42, 21) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func TestAssertGreaterThan(t *testing.T) {
	// assert.GreaterThan() should not panic
	functionUnderTest := "assert.GreaterThan"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When value A is greater than value B
	assert.GreaterThan(42, 21)
}

func TestAssertGreaterThanWithEqualValues(t *testing.T) {
	// assert.GreaterThan() should panic
	functionUnderTest := "assert.GreaterThan"
	expectedErrorMessage := "Failed asserting that value A ('42') is greater than value B ('42')"
	defer test.ShouldPanic(t, functionUnderTest, expectedErrorMessage)

	// When value A is equal to value B
	assert.GreaterThan(42, 42)
}

func TestAssertGreaterThanWithBGreaterThanA(t *testing.T) {
	// assert.GreaterThan() should panic
	functionUnderTest := "assert.GreaterThan"
	expectedErrorMessage := "Failed asserting that value A ('21') is greater than value B ('42')"
	defer test.ShouldPanic(t, functionUnderTest, expectedErrorMessage)

	// When value A is less than value B
	assert.GreaterThan(21, 42)
}
