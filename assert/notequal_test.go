package assert_test

import (
	"testing"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
)

func BenchmarkAssertNotEqual(b *testing.B) {
	codeUnderTest := func() { assert.NotEqual(42, 21) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func TestAssertNotEqual(t *testing.T) {
	// assert.NotEqual() should not panic
	functionUnderTest := "assert.NotEqual"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When value a is not equal to value b
	assert.NotEqual(42, 21)
}

func TestAssertNotEqualWithEqualValues(t *testing.T) {
	// assert.NotEqual() should panic
	functionUnderTest := "assert.NotEqual"
	expectedErrorMessage := "Failed asserting that value A and B ('Hello') are not equal"
	defer test.ShouldPanic(t, functionUnderTest, expectedErrorMessage)

	// When value a is different from value b
	assert.NotEqual("Hello", "Hello")
}
