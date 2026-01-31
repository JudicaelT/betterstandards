package assert_test

import (
	"testing"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test"
)

func TestAssertNotNil(t *testing.T) {
	// assert.NotNil() should not panic
	functionUnderTest := "assert.NotNil"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When a non-nil value is passed
	assert.NotNil("Hello world!")
}

func TestAssertNotNilWithNil(t *testing.T) {
	// assert.NotNil() should panic
	functionUnderTest := "assert.NotNil"
	expectedMessage := "Failed asserting that value is not nil"
	defer test.ShouldPanic(t, functionUnderTest, expectedMessage)

	// When a nil value is passed
	assert.NotNil(nil)
}
