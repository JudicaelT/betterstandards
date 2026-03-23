package cast_test

import (
	"encoding/binary"
	"testing"

	"github.com/JudicaelT/betterstandards/cast"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	"github.com/stretchr/testify/assert"
)

func BenchmarkCastBytesToFloat32LittleEndianWithMoreThanThreeBytes(b *testing.B) {
	bytes := []byte{0, 0, 40, 66}
	codeUnderTest := func() { cast.BytesToFloat32(bytes, binary.LittleEndian) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkCastBytesToFloat32BigEndianWithMoreThanThreeBytes(b *testing.B) {
	bytes := []byte{66, 40, 0}
	codeUnderTest := func() { cast.BytesToFloat32(bytes, binary.BigEndian) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkCastBytesToFloat32WithLessThanFourBytes(b *testing.B) {
	bytes := []byte{66, 40, 0, 0}
	codeUnderTest := func() { cast.BytesToFloat32(bytes, binary.LittleEndian) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func TestCastBytesToFloat32LittleEndianWithMoreThanThreeBytes(t *testing.T) {
	// Given an slice of four bytes and an slice of more than four bytes
	fourByteSlice := []byte{0, 0, 40, 66}
	fiveByteSlice := []byte{0, 0, 40, 66, 200}

	// When we cast the slices to float32 using LittleEndian
	for _, bytes := range [][]byte{fourByteSlice, fiveByteSlice} {
		result, err := cast.BytesToFloat32(bytes, binary.LittleEndian)

		// Then it should return
		// math.Float32frombits(66 * 256^3 + 40 * 256^2)
		// = math.Float32frombits(0x42280000) = 42
		assert.Equal(
			t,
			float32(42),
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

func TestCastBytesToFloat32BigEndianWithMoreThanThreeByte(t *testing.T) {
	// Given an slice of four bytes and an slice of more than four bytes
	fourByteSlice := []byte{66, 40, 0, 0}
	fiveByteSlice := []byte{66, 40, 0, 0, 200}

	// When we cast the slices to float32 using BigEndian
	for _, bytes := range [][]byte{fourByteSlice, fiveByteSlice} {
		result, err := cast.BytesToFloat32(bytes, binary.BigEndian)

		// Then it should return
		// math.Float32frombits(66 * 256^3 + 40 * 256^2)
		// = math.Float32frombits(0x42280000) = 42
		assert.Equal(
			t,
			float32(42),
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

func TestCastBytesToFloat32WithLessThanFourBytes(t *testing.T) {
	// Given an slice of three bytes
	threeByteSlice := []byte{66, 40, 0}

	// When we cast the slice to an float32
	result, err := cast.BytesToFloat32(threeByteSlice, binary.LittleEndian)

	// Then it should return 0
	assert.Equal(t, float32(0), result)

	// And also return an error with the expected message
	assert.Error(t, err)
	expectedMessage := "Cannot convert []byte to float32 because it contains less than 4 bytes (3 given)"
	actualMessage := err.Error()
	assert.Equal(t, expectedMessage, actualMessage)
}
