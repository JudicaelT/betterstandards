package cast_test

import (
	"encoding/binary"
	"testing"

	"github.com/JudicaelT/betterstandards/cast"
	"github.com/JudicaelT/betterstandards/internal/test/benchmark"
	"github.com/stretchr/testify/assert"
)

func BenchmarkCastBytesToFloat64LittleEndianWithMoreThanSevenBytes(b *testing.B) {
	bytes := []byte{0, 0, 0, 0, 0, 0, 69, 64}
	codeUnderTest := func() { cast.BytesToFloat64(bytes, binary.LittleEndian) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkCastBytesToFloat64BigEndianWithMoreThanSevenBytes(b *testing.B) {
	bytes := []byte{64, 69, 0, 0, 0, 0, 0, 0}
	codeUnderTest := func() { cast.BytesToFloat64(bytes, binary.BigEndian) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func BenchmarkCastBytesToFloat64WithLessThanHeightBytes(b *testing.B) {
	bytes := []byte{0, 0, 0, 0, 0, 0, 69}
	codeUnderTest := func() { cast.BytesToFloat64(bytes, binary.LittleEndian) }
	benchmark.AvgRuntime(b, codeUnderTest)
	benchmark.AssertNoAllocs(b, codeUnderTest)
}

func TestCastBytesToFloat64LittleEndianWithMoreThanSevenBytes(t *testing.T) {
	// Given an slice of height bytes and an slice of more than height bytes
	heightByteSlice := []byte{0, 0, 0, 0, 0, 0, 69, 64}
	nineByteSlice := []byte{0, 0, 0, 0, 0, 0, 69, 64, 40}

	// When we cast the slices to float64 using LittleEndian
	for _, bytes := range [][]byte{heightByteSlice, nineByteSlice} {
		result, err := cast.BytesToFloat64(bytes, binary.LittleEndian)

		// Then it should return
		// math.Float64frombits(64 * 256^7 + 69 * 256^6)
		// = math.Float64frombits(0x4045000000000000) = 42
		assert.Equal(
			t,
			float64(42),
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

func TestCastBytesToFloat64BigEndianWithMoreThanSevenBytes(t *testing.T) {
	// Given an slice of height bytes and an slice of more than height bytes
	heightByteSlice := []byte{64, 69, 0, 0, 0, 0, 0, 0}
	nineByteSlice := []byte{64, 69, 0, 0, 0, 0, 0, 0, 64}

	// When we cast the slices to float64 using BigEndian
	for _, bytes := range [][]byte{heightByteSlice, nineByteSlice} {
		result, err := cast.BytesToFloat64(bytes, binary.BigEndian)

		// Then it should return
		// math.Float64frombits(64 * 256^7 + 69 * 256^6)
		// = math.Float64frombits(0x4045000000000000) = 42
		assert.Equal(
			t,
			float64(42),
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

func TestCastBytesToFloat64WithLessThanHeightBytes(t *testing.T) {
	// Given an slice of seven bytes
	sevenByteSlice := []byte{0, 0, 0, 0, 0, 0, 69}

	// When we cast the slice to an float64
	result, err := cast.BytesToFloat64(sevenByteSlice, binary.LittleEndian)

	// Then it should return 0
	assert.Equal(t, float64(0), result)

	// And also return an error with the expected message
	assert.Error(t, err)
	expectedMessage := "Cannot convert []byte to float64 because it contains less than 8 bytes (7 given)"
	actualMessage := err.Error()
	assert.Equal(t, expectedMessage, actualMessage)
}
