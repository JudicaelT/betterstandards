package assert_test

import (
	"testing"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
)

func BenchmarkAssertSliceLen(b *testing.B) {
	slice := []int{1, 2, 3}
	codeUnderTest := func() { assert.SliceLen(3, slice) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkAssertStringLen(b *testing.B) {
	str := "Hello world"
	codeUnderTest := func() { assert.StringLen(11, str) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkAssertMapLen(b *testing.B) {
	m := make(map[int]int)
	m[42] = 42
	m[21] = 21
	codeUnderTest := func() { assert.MapLen(2, m) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkAssertChannelLen(b *testing.B) {
	channel := make(chan int, 2)
	channel <- 42
	channel <- 21
	codeUnderTest := func() { assert.ChannelLen(2, channel) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func TestAssertSliceLen(t *testing.T) {
	// assert.SliceLen() should not panic
	functionUnderTest := "assert.SliceLen"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When the length of the slice matches
	assert.SliceLen(3, []int{1, 2, 3})
}

func TestAssertSliceLenWithDifferentLen(t *testing.T) {
	// assert.SliceLen() should panic
	functionUnderTest := "assert.SliceLen"
	expectedMessage := "Failed asserting that length of slice is 3. Got 2"
	defer test.ShouldPanic(t, functionUnderTest, expectedMessage)

	// When the length of the slice does not match
	assert.SliceLen(3, []int{1, 2})
}

func TestAssertStringLen(t *testing.T) {
	// assert.StringLen() should not panic
	functionUnderTest := "assert.StringLen"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When the length of the string matches
	assert.StringLen(11, "Hello world")
}

func TestAssertStringLenWithDifferentLen(t *testing.T) {
	// assert.StringLen() should panic
	functionUnderTest := "assert.StringLen"
	expectedMessage := "Failed asserting that length of string is 3. Got 11"
	defer test.ShouldPanic(t, functionUnderTest, expectedMessage)

	// When the length of the string does not match
	assert.StringLen(3, "Hello world")
}

func TestAssertMapLen(t *testing.T) {
	// assert.MapLen() should not panic
	functionUnderTest := "assert.MapLen"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When the length of the map matches
	m := make(map[int]int)
	m[42] = 42
	assert.MapLen(1, m)
}

func TestAssertMapLenWithDifferentLen(t *testing.T) {
	// assert.MapLen() should panic
	functionUnderTest := "assert.MapLen"
	expectedMessage := "Failed asserting that length of map is 1. Got 2"
	defer test.ShouldPanic(t, functionUnderTest, expectedMessage)

	// When the length of the map does not match
	m := make(map[int]int)
	m[42] = 42
	m[21] = 21
	assert.MapLen(1, m)
}

func TestAssertChannelLen(t *testing.T) {
	// assert.ChannelLen() should not panic
	functionUnderTest := "assert.ChannelLen"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When the length of the channel matches
	channel := make(chan int, 2)
	channel <- 42
	channel <- 21
	assert.ChannelLen(2, channel)
}

func TestAssertChannelLenWithDifferentLen(t *testing.T) {
	// assert.ChannelLen() should panic
	functionUnderTest := "assert.ChannelLen"
	expectedMessage := "Failed asserting that length of channel is 2. Got 1"
	defer test.ShouldPanic(t, functionUnderTest, expectedMessage)

	// When the length of the channel does not match
	channel := make(chan int, 2)
	channel <- 42
	assert.ChannelLen(2, channel)
}
