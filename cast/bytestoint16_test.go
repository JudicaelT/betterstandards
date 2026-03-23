package cast_test

import (
	"encoding/binary"
	"testing"

	"github.com/JudicaelT/betterstandards/cast"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	"github.com/stretchr/testify/assert"
)

func BenchmarkCastBytesToInt16LittleEndianWithMoreThanOneByte(b *testing.B) {
	bytes := []byte{255, 42}
	codeUnderTest := func() { cast.BytesToInt16(bytes, binary.LittleEndian) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkCastBytesToInt16BigEndianWithMoreThanOneByte(b *testing.B) {
	bytes := []byte{255, 42}
	codeUnderTest := func() { cast.BytesToInt16(bytes, binary.BigEndian) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkCastBytesToInt16WithLessThanTwoBytes(b *testing.B) {
	bytes := []byte{255}
	codeUnderTest := func() { cast.BytesToInt16(bytes, binary.LittleEndian) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func TestCastBytesToInt16LittleEndianWithMoreThanOneByte(t *testing.T) {
	// Given a slice of two bytes and a slice of more than two bytes
	twoByteSlice := []byte{255, 42}
	threeByteSlice := []byte{255, 42, 33, 200, 42}

	// When we cast the slices to int16 LittleEndian
	for _, bytes := range [][]byte{twoByteSlice, threeByteSlice} {
		result, err := cast.BytesToInt16(bytes, binary.LittleEndian)

		// Then it should return 42 * 256 + 255 = 11007
		assert.Equal(
			t,
			int16(11007),
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

func TestCastBytesToInt16BigEndianWithMoreThanOneByte(t *testing.T) {
	// Given a slice of two bytes and a slice of more than two bytes
	twoByteSlice := []byte{255, 42}
	threeByteSlice := []byte{255, 42, 33, 200, 42}

	// When we cast the slices to int16 using BigEndian
	for _, bytes := range [][]byte{twoByteSlice, threeByteSlice} {
		result, err := cast.BytesToInt16(bytes, binary.BigEndian)

		// Then it should return 255 * 256 + 42 - 65536 = -214
		assert.Equal(
			t,
			int16(-214),
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

func TestCastBytesToInt16WithLessThanTwoBytes(t *testing.T) {
	// Given a slice of one byte
	oneByteSlice := []byte{255}

	// When we cast the slice to an int16
	result, err := cast.BytesToInt16(oneByteSlice, binary.LittleEndian)

	// Then it should return 0
	assert.Equal(t, int16(0), result)
	// And also return an error with the expected message
	assert.Error(t, err)
	expectedMessage := "Cannot convert []byte to int16 because it contains less than 2 bytes (1 given)"
	actualMessage := err.Error()
	assert.Equal(t, expectedMessage, actualMessage)
}
