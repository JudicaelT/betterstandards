package assert_test

import (
	"errors"
	"testing"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test"
)

func TestAssertNoError(t *testing.T) {
	// assert.NoError() should not panic
	functionUnderTest := "assert.NoError"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When no error is passed
	assert.NoError(nil)
}

func TestAssertNoErrorWithError(t *testing.T) {
	// Given the following error
	errorMessage := "Goodbye world!"
	err := errors.New(errorMessage)

	// assert.NoError() should panic
	functionUnderTest := "assert.NoError"
	defer test.ShouldPanic(t, functionUnderTest, errorMessage)

	// When an error is passed
	assert.NoError(err)
}
