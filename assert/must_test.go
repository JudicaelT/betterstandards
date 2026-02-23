package assert_test

import (
	"errors"
	"testing"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
)

func BenchmarkAssertMust(b *testing.B) {
	codeUnderTest := func() { assert.Must(42, nil) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func TestAssertMust(t *testing.T) {
	// Given a value and no error
	var value int8 = 42
	var err error = nil

	// assert.Nil() should not panic
	functionUnderTest := "assert.Must"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When no error is passed
	var actual int8 = assert.Must(value, err)

	// And it should return the value that was passed
	if value != actual {
		t.Errorf(
			"%s did not return the given value %d, returned %d instead",
			functionUnderTest,
			value,
			actual,
		)
	}
}

func TestAssertMustWithError(t *testing.T) {
	// Given a value and an error
	errMessage := "Expected message"
	var value int8 = 42
	var err error = errors.New(errMessage)

	// assert.Must() should panic
	functionUnderTest := "assert.Must"
	defer test.ShouldPanic(t, functionUnderTest, errMessage)

	// When an error is passed
	assert.Must(value, err)
}
