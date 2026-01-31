package assert_test

import (
	"testing"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test"
)

func TestAssertEqual(t *testing.T) {
	// assert.Equal() should not panic
	functionUnderTest := "assert.Equal"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When value a is equal to value b
	assert.Equal(42, 42)
}

func TestAssertEqualWithDifferentValues(t *testing.T) {
	// assert.Equal() should panic

	functionUnderTest := "assert.Equal"
	expectedErrorMessage := "Failed asserting that 'Hello' is equal to 'World'"
	defer test.ShouldPanic(t, functionUnderTest, expectedErrorMessage)

	// When value a is different from value b
	assert.Equal("Hello", "World")
}
