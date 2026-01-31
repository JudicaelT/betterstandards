package assert_test

import (
	"testing"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test"
)

func TestAssertFalse(t *testing.T) {
	// assert.False() should not panic
	functionUnderTest := "assert.False"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When false is passed
	assert.False(false)
}

func TestAssertFalseWithFalse(t *testing.T) {
	// assert.False() should panic
	functionUnderTest := "assert.False"
	expectedMessage := "Failed asserting that value is false"
	defer test.ShouldPanic(t, functionUnderTest, expectedMessage)

	// When true is passed
	assert.False(true)
}
