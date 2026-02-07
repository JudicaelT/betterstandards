package cast_test

import (
	"encoding/binary"
	"testing"

	"github.com/JudicaelT/betterstandards/cast"
)

func TestCastBytesToFloat32LittleEndianWithMoreThanThreeBytes(t *testing.T) {
	// Given an slice of four bytes and an slice of more than four bytes
	fourByteSlice := []byte{0, 0, 40, 66}
	fiveByteSlice := []byte{0, 0, 40, 66, 200}

	// When we cast the slices to float32 using LittleEndian
	for _, byteSlice := range [][]byte{fourByteSlice, fiveByteSlice} {
		result, err := cast.BytesToFloat32(byteSlice, binary.LittleEndian)

		// Then it should return
		// math.Float32frombits(66 * 256^3 + 40 * 256^2)
		// = math.Float32frombits(0x42280000) = 42
		if result != 42 {
			t.Errorf(
				"(Dataset: %v) Failed asserting that result is equal to 42. Got %f",
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

func TestCastBytesToFloat32BigEndianWithMoreThanThreeByte(t *testing.T) {
	// Given an slice of four bytes and an slice of more than four bytes
	fourByteSlice := []byte{66, 40, 0, 0}
	fiveByteSlice := []byte{66, 40, 0, 0, 200}

	// When we cast the slices to float32 using BigEndian
	for _, byteSlice := range [][]byte{fourByteSlice, fiveByteSlice} {
		result, err := cast.BytesToFloat32(byteSlice, binary.BigEndian)

		// Then it should return
		// math.Float32frombits(66 * 256^3 + 40 * 256^2)
		// = math.Float32frombits(0x42280000) = 42
		if result != 42 {
			t.Errorf(
				"(Dataset: %v) Failed asserting that result is equal to 42 Got %f",
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

func TestCastBytesToFloat32WithLessThanFourBytes(t *testing.T) {
	// Given an slice of three bytes
	threeByteSlice := []byte{66, 40, 0}

	// When we cast the slice to an float32
	result, err := cast.BytesToFloat32(threeByteSlice, binary.LittleEndian)

	// Then it should return 0
	if result != 0 {
		t.Error("Failed asserting that result is equal to 0. Got", result)
	}

	// And also return an error with the expected message
	if err == nil {
		t.Fatalf("Failed asserting that error is not nil")
	}
	expectedMessage := "Cannot convert []byte to float32 because it contains less than 4 bytes (3 given)"
	actualMessage := err.Error()
	if expectedMessage != actualMessage {
		t.Errorf(
			"Expected error message '%s' but got '%s'",
			expectedMessage,
			actualMessage,
		)
	}
}
