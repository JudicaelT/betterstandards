package assert_test

import (
	"testing"

	"github.com/JudicaelT/betterstandards/assert"
	"github.com/JudicaelT/betterstandards/internal/test"
)

func TestAssertEmptyWithSlice(t *testing.T) {
	// Given an empty slice
	emptySlice := []any{}

	// assert.EmptySlice() should not panic
	functionUnderTest := "assert.EmptySlice"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When the empty slice is passed
	assert.EmptySlice(emptySlice)
}

func TestAssertEmptyWithNonEmptySlice(t *testing.T) {
	// Given a non-empty slice
	nonEmptySlice := []int8{1, 2, 3}

	// assert.EmptySlice() should panic
	functionUnderTest := "assert.EmptySlice"
	expectedMessage := "Failed asserting that slice is empty. Found '3' elements"
	defer test.ShouldPanic(t, functionUnderTest, expectedMessage)

	// When the non-empty slice is passed
	assert.EmptySlice(nonEmptySlice)
}

func TestAssertEmptyWithString(t *testing.T) {
	// Given an empty string
	emptyString := ""

	// assert.EmptyString() should not panic
	functionUnderTest := "assert.EmptyString"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When the empty string is passed
	assert.EmptyString(emptyString)
}

func TestAssertEmptyWithNonEmptyString(t *testing.T) {
	// Given a non-empty string
	nonEmptyString := "non-empty-string"

	// assert.EmptyString() should panic
	functionUnderTest := "assert.EmptyString"
	expectedMessage := "Failed asserting that string is empty. Got 'non-empty-string'"
	defer test.ShouldPanic(t, functionUnderTest, expectedMessage)

	// When the non-empty string is passed
	assert.EmptyString(nonEmptyString)
}

func TestAssertEmptyWithMap(t *testing.T) {
	// Given an empty map
	emptyMap := make(map[int]string)

	// assert.EmptyMap() should not panic
	functionUnderTest := "assert.EmptyMap"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When the empty map is passed
	assert.EmptyMap(emptyMap)
}

func TestAssertEmptyWithNonEmptyMap(t *testing.T) {
	// Given a non-empty map
	nonEmptyMap := make(map[int]int)
	nonEmptyMap[42] = 42

	// assert.EmptyMap() should panic
	functionUnderTest := "assert.EmptyMap"
	expectedMessage := "Failed asserting that map is empty. Found '1' elements"
	defer test.ShouldPanic(t, functionUnderTest, expectedMessage)

	// When the non-empty map is passed
	assert.EmptyMap(nonEmptyMap)
}

func TestAssertEmptyWithChannel(t *testing.T) {
	// Given an empty channel
	emptyChannel := make(chan any)

	// assert.EmptyChannel() should not panic
	functionUnderTest := "assert.EmptyChannel"
	defer test.ShouldNotPanic(t, functionUnderTest)

	// When the empty channel is passed
	assert.EmptyChannel(emptyChannel)
}

func TestAssertEmptyWithNonEmptyChannel(t *testing.T) {
	// Given a non-empty channel
	nonEmptyChannel := make(chan int, 1)
	nonEmptyChannel <- 1

	// assert.EmptyChannel() should not panic
	functionUnderTest := "assert.EmptyChannel"
	expectedMessage := "Failed asserting that channel is empty. Found '1' elements"
	defer test.ShouldPanic(t, functionUnderTest, expectedMessage)

	// When the non-empty channel is passed
	assert.EmptyChannel(nonEmptyChannel)
}
