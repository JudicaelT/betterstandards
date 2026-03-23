package cast_test

import (
	"encoding/binary"
	"testing"

	"github.com/JudicaelT/betterstandards/cast"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	"github.com/stretchr/testify/assert"
)

func BenchmarkCastBytesToInt32LittleEndianWithMoreThanThreeBytes(b *testing.B) {
	bytes := []byte{255, 42, 33, 20}
	codeUnderTest := func() { cast.BytesToInt32(bytes, binary.LittleEndian) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkCastBytesToInt32BigEndianWithMoreThanThreeBytes(b *testing.B) {
	bytes := []byte{255, 42, 33, 20}
	codeUnderTest := func() { cast.BytesToInt32(bytes, binary.BigEndian) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkCastBytesToInt32WithLessThanFourBytes(b *testing.B) {
	bytes := []byte{255, 42, 33}
	codeUnderTest := func() { cast.BytesToInt32(bytes, binary.LittleEndian) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func TestCastBytesToInt32LittleEndianWithMoreThanThreeBytes(t *testing.T) {
	// Given a slice of four bytes and a slice of more than four bytes
	fourByteSlice := []byte{255, 42, 33, 20}
	fiveByteSlice := []byte{255, 42, 33, 20, 42}

	// When we cast the slices to int32 using LittleEndian
	for _, bytes := range [][]byte{fourByteSlice, fiveByteSlice} {
		result, err := cast.BytesToInt32(bytes, binary.LittleEndian)

		// Then it should return
		// 20 * 256^3 + 33 * 256^2 + 42 * 256 + 255 = 337718015
		assert.Equal(
			t,
			int32(337718015),
			result,
			"Test failed for dataset: %v",
			bytes,
		)
		// And not return an error
		assert.NoError(
			t,
			err,
			"Test failed for dataset: %v",
			bytes,
		)
	}
}

func TestCastBytesToInt32BigEndianWithMoreThanThreeByte(t *testing.T) {
	// Given a slice of four bytes and a slice of more than four bytes
	fourByteSlice := []byte{255, 42, 33, 20}
	fiveByteSlice := []byte{255, 42, 33, 20, 42}

	// When we cast the slices to int32 using BigEndian
	for _, bytes := range [][]byte{fourByteSlice, fiveByteSlice} {
		result, err := cast.BytesToInt32(bytes, binary.BigEndian)

		// Then it should return
		// 255 * 256^3 + 42 * 256^2 + 33 * 256 + 20 - 4294967296 = -14016236
		assert.Equal(
			t,
			int32(-14016236),
			result,
			"Test failed for dataset: %v",
			bytes,
		)
		// And not return an error
		assert.NoError(
			t,
			err,
			"Test failed for dataset: %v",
			bytes,
		)
	}
}

func TestCastBytesToInt32WithLessThanFourBytes(t *testing.T) {
	// Given a slice of three bytes
	threeByteSlice := []byte{255, 42, 33}

	// When we cast the slice to an int32
	result, err := cast.BytesToInt32(threeByteSlice, binary.LittleEndian)

	// Then it should return 0
	assert.Equal(t, int32(0), result)
	// And also return an error with the expected message
	assert.Error(t, err)
	expectedMessage := "Cannot convert []byte to int32 because it contains less than 4 bytes (3 given)"
	actualMessage := err.Error()
	assert.Equal(t, expectedMessage, actualMessage)
}
