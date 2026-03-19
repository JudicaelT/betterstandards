package assert_test

import (
	"testing"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
)

func BenchmarkAssertLessThan(b *testing.B) {
	codeUnderTest := func() { assert.LessThan(21, 42) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func TestAssertLessThan(t *testing.T) {
	// assert.LessThan() should not panic
	functionUnderTest := "assert.LessThan"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When value A is less than value B
	assert.LessThan(21, 42)
}

func TestAssertLessThanWithEqualValues(t *testing.T) {
	// assert.LessThan() should panic
	functionUnderTest := "assert.LessThan"
	expectedErrorMessage := "Failed asserting that value A ('42') is less than value B ('42')"
	defer test.ShouldPanic(t, functionUnderTest, expectedErrorMessage)

	// When value A is equal to value B
	assert.LessThan(42, 42)
}

func TestAssertLessThanWithBLessThanA(t *testing.T) {
	// assert.LessThan() should panic
	functionUnderTest := "assert.LessThan"
	expectedErrorMessage := "Failed asserting that value A ('42') is less than value B ('21')"
	defer test.ShouldPanic(t, functionUnderTest, expectedErrorMessage)

	// When value A is greater than value B
	assert.LessThan(42, 21)
}
