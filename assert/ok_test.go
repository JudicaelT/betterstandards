package assert_test

import (
	"testing"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	testify "github.com/stretchr/testify/assert"
)

func BenchmarkAssertOk(b *testing.B) {
	okFunc := func() (int, bool) { return 42, true }
	codeUnderTest := func() { assert.Ok(okFunc()) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func TestAssertOk(t *testing.T) {
	// Given a function that returns true as its second return value
	expected := "test"
	okFunc := func() (string, bool) { return expected, true }

	// assert.Ok should not panic when we pass the function
	var actual string
	testify.NotPanics(t, func() { actual = assert.Ok(okFunc()) })
	// And it should return okFunc's first value
	testify.Equal(t, expected, actual)
}

func TestAssertOkWithFalse(t *testing.T) {
	// Given a function that returns false as its second return value
	nokFunc := func() (int, bool) { return 42, false }

	// assert.Ok should panic when we pass the function
	testify.Panics(t, func() { assert.Ok(nokFunc()) })
}
