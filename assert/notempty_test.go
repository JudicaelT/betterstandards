package assert_test

import (
	"testing"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
)

func BenchmarkAssertNotEmptySlice(b *testing.B) {
	slice := []any{42}
	codeUnderTest := func() { assert.NotEmptySlice(slice) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkAssertNotEmptyString(b *testing.B) {
	str := "Hello world"
	codeUnderTest := func() { assert.NotEmptyString(str) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkAssertNotEmptyMap(b *testing.B) {
	m := make(map[int]string)
	m[0] = "Hello world"
	codeUnderTest := func() { assert.NotEmptyMap(m) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkAssertNotEmptyChannel(b *testing.B) {
	channel := make(chan any, 1)
	channel <- "Hello world"
	codeUnderTest := func() { assert.NotEmptyChannel(channel) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func TestAssertNotEmptyWithSlice(t *testing.T) {
	// Given a non-empty slice
	slice := []any{42}

	// assert.NotEmptySlice() should not panic
	functionUnderTest := "assert.NotEmptySlice"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When the non-empty slice is passed
	assert.NotEmptySlice(slice)
}

func TestAssertNotEmptyWithEmptySlice(t *testing.T) {
	// Given an empty slice
	emptySlice := []int8{}

	// assert.NotEmptySlice() should panic
	functionUnderTest := "assert.NotEmptySlice"
	expectedMessage := "Failed asserting that the given slice is not empty"
	defer test.ShouldPanic(t, functionUnderTest, expectedMessage)

	// When the empty slice is passed
	assert.NotEmptySlice(emptySlice)
}

func TestAssertNotEmptyWithString(t *testing.T) {
	// Given a non-empty string
	str := "Hello world"

	// assert.NotEmptyString() should not panic
	functionUnderTest := "assert.NotEmptyString"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When the non-empty string is passed
	assert.NotEmptyString(str)
}

func TestAssertNotEmptyWithEmptyString(t *testing.T) {
	// Given an empty string
	emptyStr := ""

	// assert.NotEmptyString() should panic
	functionUnderTest := "assert.NotEmptyString"
	expectedMessage := "Failed asserting that the given string is not empty"
	defer test.ShouldPanic(t, functionUnderTest, expectedMessage)

	// When the empty string is passed
	assert.NotEmptyString(emptyStr)
}

func TestAssertNotEmptyWithMap(t *testing.T) {
	// Given a non-empty map
	m := make(map[int]string)
	m[0] = "Hello world"

	// assert.NotEmptyMap() should not panic
	functionUnderTest := "assert.NotEmptyMap"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When the non-empty map is passed
	assert.NotEmptyMap(m)
}

func TestAssertNotEmptyWithEmptyMap(t *testing.T) {
	// Given an empty map
	emptyMap := make(map[int]int)

	// assert.NotEmptyMap() should panic
	functionUnderTest := "assert.NotEmptyMap"
	expectedMessage := "Failed asserting that the given map is not empty"
	defer test.ShouldPanic(t, functionUnderTest, expectedMessage)

	// When the empty map is passed
	assert.NotEmptyMap(emptyMap)
}

func TestAssertNotEmptyWithChannel(t *testing.T) {
	// Given a non-empty channel
	channel := make(chan any, 1)
	channel <- "Hello world"

	// assert.NotEmptyChannel() should not panic
	functionUnderTest := "assert.NotEmptyChannel"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When the non-empty channel is passed
	assert.NotEmptyChannel(channel)
}

func TestAssertNotEmptyWithEmptyChannel(t *testing.T) {
	// Given an empty channel
	emptyChannel := make(chan int, 1)

	// assert.NotEmptyChannel() should panic
	functionUnderTest := "assert.NotEmptyChannel"
	expectedMessage := "Failed asserting that the given channel is not empty"
	defer test.ShouldPanic(t, functionUnderTest, expectedMessage)

	// When the empty channel is passed
	assert.NotEmptyChannel(emptyChannel)
}
