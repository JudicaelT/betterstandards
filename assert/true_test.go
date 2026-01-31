package assert_test

import (
	"testing"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test"
)

func TestAssertTrue(t *testing.T) {
	// assert.True() should not panic
	functionUnderTest := "assert.True"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When true is passed
	assert.True(true)
}

func TestAssertTrueWithFalse(t *testing.T) {
	// assert.True() should panic
	functionUnderTest := "assert.True"
	expectedMessage := "Failed asserting that value is true"
	defer test.ShouldPanic(t, functionUnderTest, expectedMessage)

	// When false is passed
	assert.True(false)
}
