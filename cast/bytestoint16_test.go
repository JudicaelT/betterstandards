package cast_test

import (
	"encoding/binary"
	"testing"

	"github.com/JudicaelT/betterstandards/cast"
)

func TestCastBytesToInt16LittleEndianWithMoreThanOneByte(t *testing.T) {
	// Given a slice of two bytes and a slice of more than two bytes
	twoByteSlice := []byte{255, 42}
	threeByteSlice := []byte{255, 42, 33, 200, 42}

	// When we cast the slices to int16 LittleEndian
	for _, byteSlice := range [][]byte{twoByteSlice, threeByteSlice} {
		result, err := cast.BytesToInt16(byteSlice, binary.LittleEndian)

		// Then it should return 42 * 256 + 255 = 11007
		if result != 11007 {
			t.Errorf(
				"(Dataset: %v) Failed asserting that result is equal to 11007. Got %d",
				byteSlice,
				result,
			)
		}

		// And not return an error
		if err != nil {
			t.Errorf(
				"(Dataset: %v) Failed asserting that error is nil: %s",
				byteSlice,
				err,
			)
		}
	}
}

func TestCastBytesToInt16BigEndianWithMoreThanOneByte(t *testing.T) {
	// Given a slice of two bytes and a slice of more than two bytes
	twoByteSlice := []byte{255, 42}
	threeByteSlice := []byte{255, 42, 33, 200, 42}

	// When we cast the slices to int16 using BigEndian
	for _, byteSlice := range [][]byte{twoByteSlice, threeByteSlice} {
		result, err := cast.BytesToInt16(byteSlice, binary.BigEndian)

		// Then it should return 255 * 256 + 42 - 65536 = -214
		if result != -214 {
			t.Errorf(
				"(Dataset: %v) Failed asserting that result is equal to -214. Got %d",
				byteSlice,
				result,
			)
		}

		// And not return an error
		if err != nil {
			t.Errorf(
				"(Dataset: %v) Failed asserting that error nil: %s",
				byteSlice,
				err,
			)
		}
	}
}

func TestCastBytesToInt16WithLessThanTwoBytes(t *testing.T) {
	// Given a slice of one byte
	oneByteSlice := []byte{255}

	// When we cast the slice to an int16
	result, err := cast.BytesToInt16(oneByteSlice, binary.LittleEndian)

	// Then it should return 0
	if result != 0 {
		t.Error("Failed asserting that result is equal to 0. Got", result)
	}

	// And also return an error with the expected message
	if err == nil {
		t.Error("Failed asserting that error is not nil")
		return
	}
	expectedMessage := "Cannot convert []byte to int16 because it contains less than 2 bytes (1 given)"
	actualMessage := err.Error()
	if expectedMessage != actualMessage {
		t.Errorf(
			"Expected error message '%s' but got '%s'",
			expectedMessage,
			actualMessage,
		)
	}
}
