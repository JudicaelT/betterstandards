package cast_test

import (
	"encoding/binary"
	"testing"

	"github.com/JudicaelT/betterstandards/cast"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	"github.com/stretchr/testify/assert"
)

func BenchmarkCastBytesToInt64LittleEndianWithMoreThanSevenBytes(b *testing.B) {
	bytes := []byte{255, 42, 33, 20, 100, 77, 234, 10}
	codeUnderTest := func() { cast.BytesToInt64(bytes, binary.LittleEndian) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkCastBytesToInt64BigEndianWithMoreThanSevenBytes(b *testing.B) {
	bytes := []byte{255, 42, 33, 20, 100, 77, 234, 10}
	codeUnderTest := func() { cast.BytesToInt64(bytes, binary.BigEndian) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkCastBytesToInt64WithLessThanHeightBytes(b *testing.B) {
	bytes := []byte{255, 42, 33, 20, 100, 77, 234}
	codeUnderTest := func() { cast.BytesToInt64(bytes, binary.LittleEndian) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func TestCastBytesToInt64LittleEndianWithMoreThanSevenBytes(t *testing.T) {
	// Given a slice of height bytes and a slice of more than height bytes
	heightByteSlice := []byte{255, 42, 33, 20, 100, 77, 234, 10}
	nineByteSlice := []byte{255, 42, 33, 20, 100, 77, 234, 10, 20}

	// When we cast the slices to int64 using LittleEndian
	for _, bytes := range [][]byte{heightByteSlice, nineByteSlice} {
		result, err := cast.BytesToInt64(bytes, binary.LittleEndian)

		// Then it should return
		// 10 * 256^7 + 234 * 256^6 + ... + 33 * 256^2 + 42 * 256 + 255 = 786526177159359231
		assert.Equal(
			t,
			int64(786526177159359231),
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

func TestCastBytesToInt64BigEndianWithMoreThanSevenBytes(t *testing.T) {
	// Given a slice of height bytes and a slice of more than height bytes
	heightByteSlice := []byte{255, 42, 33, 20, 100, 77, 234, 10}
	nineByteSlice := []byte{255, 42, 33, 20, 100, 77, 234, 10, 20}

	// When we cast the slices to int64 using LittleEndian
	for _, bytes := range [][]byte{heightByteSlice, nineByteSlice} {
		result, err := cast.BytesToInt64(bytes, binary.BigEndian)

		// Then it should return
		// 255 * 256^7 + 42 * 256^6 + ... + 77 * 256^2 + 234 * 256 + 10 = -60199273550190070
		assert.Equal(
			t,
			int64(-60199273550190070),
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

func TestCastBytesToInt64WithLessThanHeightBytes(t *testing.T) {
	// Given a slice of seven bytes
	sevenByteSlice := []byte{255, 42, 33, 20, 100, 77, 234}

	// When we cast the slice to an int64
	result, err := cast.BytesToInt64(sevenByteSlice, binary.LittleEndian)

	// Then it should return 0
	assert.Equal(t, int64(0), result)
	// And also return an error with the expected message
	assert.Error(t, err)
	expectedMessage := "Cannot convert []byte to int64 because it contains less than 8 bytes (7 given)"
	actualMessage := err.Error()
	assert.Equal(t, expectedMessage, actualMessage)
}
