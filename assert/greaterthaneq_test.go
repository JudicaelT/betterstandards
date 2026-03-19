package assert_test

import (
	"testing"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
)

func BenchmarkAssertGreaterThanEq(b *testing.B) {
	codeUnderTest := func() { assert.GreaterThanEq(42, 21) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func TestAssertGreaterThanEq(t *testing.T) {
	// assert.GreaterThanEq() should not panic
	functionUnderTest := "assert.GreaterThanEq"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When value A is greater than value B
	assert.GreaterThanEq(42, 21)
}

func TestAssertGreaterThanEqWithEqualValues(t *testing.T) {
	// assert.GreaterThanEq() should not panic
	functionUnderTest := "assert.GreaterThanEq"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When value A is equal to value B
	assert.GreaterThanEq(42, 42)
}

func TestAssertGreaterThanEqWithBGreaterThanEqA(t *testing.T) {
	// assert.GreaterThanEq() should panic
	functionUnderTest := "assert.GreaterThanEq"
	expectedErrorMessage := "Failed asserting that value A ('21') is equal or greater than value B ('42')"
	defer test.ShouldPanic(t, functionUnderTest, expectedErrorMessage)

	// When value A is less than value B
	assert.GreaterThanEq(21, 42)
}
