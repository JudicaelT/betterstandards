package cast_test

import (
	"encoding/binary"
	"testing"

	"github.com/JudicaelT/betterstandards/cast"
)

func TestCastBytesToInt32LittleEndianWithMoreThanThreeBytes(t *testing.T) {
	// Given a slice of four bytes and a slice of more than four bytes
	fourByteSlice := []byte{255, 42, 33, 20}
	fiveByteSlice := []byte{255, 42, 33, 20, 42}

	// When we cast the slices to int32 using LittleEndian
	for _, byteSlice := range [][]byte{fourByteSlice, fiveByteSlice} {
		result, err := cast.BytesToInt32(byteSlice, binary.LittleEndian)

		// Then it should return
		// 20 * 256^3 + 33 * 256^2 + 42 * 256 + 255 = 337718015
		if result != 337718015 {
			t.Errorf(
				"(Dataset: %v) Failed asserting that result is equal to 337718015. Got %d",
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

func TestCastBytesToInt32BigEndianWithMoreThanThreeByte(t *testing.T) {
	// Given a slice of four bytes and a slice of more than four bytes
	fourByteSlice := []byte{255, 42, 33, 20}
	fiveByteSlice := []byte{255, 42, 33, 20, 42}

	// When we cast the slices to int32 using BigEndian
	for _, byteSlice := range [][]byte{fourByteSlice, fiveByteSlice} {
		result, err := cast.BytesToInt32(byteSlice, binary.BigEndian)

		// Then it should return
		// 255 * 256^3 + 42 * 256^2 + 33 * 256 + 20 - 4294967296 = -14016236
		if result != -14016236 {
			t.Errorf(
				"(Dataset: %v) Failed asserting that result is equal to -14016236 Got %d",
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

func TestCastBytesToInt32WithLessThanFourBytes(t *testing.T) {
	// Given a slice of three bytes
	threeByteSlice := []byte{255, 42, 33}

	// When we cast the slice to an int32
	result, err := cast.BytesToInt32(threeByteSlice, binary.LittleEndian)

	// Then it should return 0
	if result != 0 {
		t.Error("Failed asserting that result is equal to 0. Got", result)
	}

	// And also return an error with the expected message
	if err == nil {
		t.Error("Failed asserting that error is not nil")
		return
	}
	expectedMessage := "Cannot convert []byte to int32 because it contains less than 4 bytes (3 given)"
	actualMessage := err.Error()
	if expectedMessage != actualMessage {
		t.Errorf(
			"Expected error message '%s' but got '%s'",
			expectedMessage,
			actualMessage,
		)
	}
}
