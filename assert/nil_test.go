package assert_test

import (
	"errors"
	"testing"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
)

func BenchmarkAssertNil(b *testing.B) {
	codeUnderTest := func() { assert.Nil(nil) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func TestAssertNil(t *testing.T) {
	// assert.Nil() should not panic
	functionUnderTest := "assert.Nil"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When a nil value is passed
	assert.Nil(nil)
}

func TestAssertNilWithNonNil(t *testing.T) {
	// Given a value that is not nil
	value := 42

	// assert.Nil() should panic
	functionUnderTest := "assert.Nil"
	expectedMessage := "Expected value to be nil. Got 'int'"
	defer test.ShouldPanic(t, functionUnderTest, expectedMessage)

	// When a value other than nil is passed
	assert.Nil(value)
}

func TestAssertNilWithError(t *testing.T) {
	// Given a value that is an error
	value := errors.New("Goodbye world")

	// assert.Nil() should panic
	functionUnderTest := "assert.Nil"
	expectedMessage := "Expected value to be nil. Got error: '" + value.Error() + "'"
	defer test.ShouldPanic(t, functionUnderTest, expectedMessage)

	// When an error is passed
	assert.Nil(value)
}
