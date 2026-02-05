package cast_test

import (
	"encoding/binary"
	"testing"

	"github.com/JudicaelT/betterstandards/cast"
)

func TestCastBytesToInt64LittleEndianWithMoreThanSevenBytes(t *testing.T) {
	// Given a slice of height bytes and a slice of more than height bytes
	heightByteSlice := []byte{255, 42, 33, 20, 100, 77, 234, 10}
	nineByteSlice := []byte{255, 42, 33, 20, 100, 77, 234, 10, 20}

	// When we cast the slices to int64 using LittleEndian
	for _, byteSlice := range [][]byte{heightByteSlice, nineByteSlice} {
		result, err := cast.BytesToInt64(byteSlice, binary.LittleEndian)

		// Then it should return
		// 10 * 256^7 + 234 * 256^6 + ... + 33 * 256^2 + 42 * 256 + 255 = 786526177159359231
		if result != 786526177159359231 {
			t.Errorf(
				"(Dataset: %v) Failed asserting that result is equal to 786526177159359231. Got %d",
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

func TestCastBytesToInt64BigEndianWithMoreThanSevenBytes(t *testing.T) {
	// Given a slice of height bytes and a slice of more than height bytes
	heightByteSlice := []byte{255, 42, 33, 20, 100, 77, 234, 10}
	nineByteSlice := []byte{255, 42, 33, 20, 100, 77, 234, 10, 20}

	// When we cast the slices to int64 using LittleEndian
	for _, byteSlice := range [][]byte{heightByteSlice, nineByteSlice} {
		result, err := cast.BytesToInt64(byteSlice, binary.BigEndian)

		// Then it should return
		// 255 * 256^7 + 42 * 256^6 + ... + 77 * 256^2 + 234 * 256 + 10 = -60199273550190070
		if result != -60199273550190070 {
			t.Errorf(
				"(Dataset: %v) Failed asserting that result is equal to -60199273550190070. Got %d",
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

func TestCastBytesToInt64WithLessThanHeightBytes(t *testing.T) {
	// Given a slice of seven bytes
	sevenByteSlice := []byte{255, 42, 33, 20, 100, 77, 234}

	// When we cast the slice to an int64
	result, err := cast.BytesToInt64(sevenByteSlice, binary.LittleEndian)

	// Then it should return 0
	if result != 0 {
		t.Error("Failed asserting that result is equal to 0. Got", result)
	}

	// And also return an error with the expected message
	if err == nil {
		t.Error("Failed asserting that error is not nil")
		return
	}
	expectedMessage := "Cannot convert []byte to int64 because it contains less than 8 bytes (7 given)"
	actualMessage := err.Error()
	if expectedMessage != actualMessage {
		t.Errorf(
			"Expected error message '%s' but got '%s'",
			expectedMessage,
			actualMessage,
		)
	}
}
