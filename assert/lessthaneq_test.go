package assert_test

import (
	"testing"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
)

func BenchmarkAssertLessThanEq(b *testing.B) {
	codeUnderTest := func() { assert.LessThanEq(21, 42) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func TestAssertLessThanEq(t *testing.T) {
	// assert.LessThanEq() should not panic
	functionUnderTest := "assert.LessThanEq"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When value A is less than value B
	assert.LessThanEq(21, 42)
}

func TestAssertLessThanEqWithEqualValues(t *testing.T) {
	// assert.LessThanEq() should not panic
	functionUnderTest := "assert.LessThanEq"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When value A is equal to value B
	assert.LessThanEq(42, 42)
}

func TestAssertLessThanEqWithBLessThanEqA(t *testing.T) {
	// assert.LessThanEq() should panic
	functionUnderTest := "assert.LessThanEq"
	expectedErrorMessage := "Failed asserting that value A ('42') is equal or less than value B ('21')"
	defer test.ShouldPanic(t, functionUnderTest, expectedErrorMessage)

	// When value A is greater than value B
	assert.LessThanEq(42, 21)
}
